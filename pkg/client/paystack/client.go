package paystack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	client    http.Client
	baseUrl   string
	secretKey string
}

func NewClient(baseUrl, secretKey string) *Client {

	return &Client{*http.DefaultClient, baseUrl, secretKey}
}

func (cl *Client) DownloadInvoices(invoiceStatus InvoiceStatus) (*ListInvoicesResponse, error) {

	// setup request and parse param to url
	params := url.Values{}
	params.Add("status", string(invoiceStatus))
	params.Add("page", "0")

	req, err := http.NewRequest(http.MethodGet, cl.baseUrl+"/paymentrequest", strings.NewReader(params.Encode()))

	if err != nil {
		return nil, err
	}
	cl.addRequiredHeaders(req)

	response, err := cl.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got response code: %d", response.StatusCode)
	}

	var listInvoicesResponse ListInvoicesResponse

	if err = json.NewDecoder(response.Body).Decode(&listInvoicesResponse); err != nil {
		return nil, err
	}

	return &listInvoicesResponse, nil

}


func (cl *Client) GetCustomer(customerIdOrEmail string) (*GetCustomerResponse, error) {
	req, err := http.NewRequest(http.MethodGet, cl.baseUrl + "/customer/" + customerIdOrEmail, nil)

	if err != nil {
		return nil, err
	}

	cl.addRequiredHeaders(req)

	response, err := cl.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusNotFound:
		return nil, ErrCustomerNotFound
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("invalid secret key")
	case http.StatusOK:
		return getCustomer(response)
	default:
		return nil, fmt.Errorf("unknown response code: %d", response.StatusCode)
		
	}

}


func (cl *Client) CreateCustomer(ccq CreateCustomerRequest) error {
  
	body, err := json.Marshal(ccq)

	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, cl.baseUrl + "/customer", bytes.NewReader(body) )

	if err != nil {
		return err
	}

	cl.addRequiredHeaders(req)
	req.Header.Add("Content-Type", "application/json")

	response, err := cl.client.Do(req)

	if err != nil {
		return err
	}

	if response.StatusCode == http.StatusOK || response.StatusCode == http.StatusCreated {
		return nil
	}

	return fmt.Errorf("could not create customer. Invalid status : %d", response.StatusCode)
}

func (cl *Client) addRequiredHeaders(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cl.secretKey))
	req.Header.Add("Accept", "application/json")
}


func getCustomer(response *http.Response) (*GetCustomerResponse, error) {
	var customer GetCustomerResponse
	if err := json.NewDecoder(response.Body).Decode(&customer); err != nil {
		return nil, err
	}

	return &customer, nil
}
