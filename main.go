package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Amount int

type DateTime string

type lineItem struct {
	Name   string `json:"name"`
	Amount Amount `json:"amount"`
}

type tax struct {
	Name   string `json:"name"`
	Amount Amount `json:"amount"`
}

type customer struct {
	ID                       int    `json:"id"`
	FirstName                string `json:"first_name"`
	LastName                 string `json:"last_name"`
	Email                    string `json:"email"`
	Code                     string `json:"customer_code"`
	Phone                    string `json:"phone"`
	RiskAction               string `json:"risk_action"`
	InternationalFormatPhone string `json:"international_format_phone"`
	Metadata                 struct {
		CallingCode string `json:"calling_code"`
	} `json:"metadata"`
}

type Invoice struct {
	ID            int        `json:"id"`
	Domain        string     `json:"domain"`
	Amount        Amount     `json:"amount"`
	Currency      string     `json:"currency"`
	DueDate       DateTime   `json:"due_date"`
	HasInvoice    bool       `json:"has_invoice"`
	InvoiceNumber int        `json:"invoice_number"`
	Description   string     `json:"description"`
	PdfUrl        string     `json:"pdf_url"`
	LineItems     []lineItem `json:"line_items"`
	Tax           []tax      `json:"tax"`
	Customer      customer   `json:"customer"`
	RequestCode   string     `json:"request_code"`
	Status        string     `json:"status"`
	Paid          bool       `json:"paid"`
	PaidAt        DateTime   `json:"paid_at"`
	CreatedAt     DateTime   `json:"created_at"`
}
type meta struct {
	Total     int `json:"total"`
	Skipped   int `json:"skipped"`
	PerPage   int `json:"perPage"`
	Page      int `json:"page"`
	PageCount int `json:"pageCount"`
}

type listInvoicesResponse struct {
	Status  bool      `json:"status"`
	Message string    `json:"message"`
	Meta    meta      `json:"meta"`
	Data    []Invoice `json:"data"`
}

const SECRET_KEY = "sk_test_3cec88e7ece6d9f5b69f33c013b94cc9142bf161"

func downloadInvoices() {
	client := http.DefaultClient

	// setup request and parse param to url
	params := url.Values{}
	params.Add("status", "Pending")
	params.Add("page", "0")

	req, err := http.NewRequest(http.MethodGet, "https://api.paystack.co/paymentrequest", strings.NewReader(params.Encode()))

	if err != nil {
		log.Fatalf("could not create request. %v", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", SECRET_KEY))
	req.URL.Query().Add("status", "pending")

	response, err := client.Do(req)

	if err != nil {
		log.Fatalf("Failed call to PS: %v\n", err)
	}

	defer response.Body.Close()

	// we only want to deserialize if we get a successful response (200) from PS
	if response.StatusCode == http.StatusOK {

		var lstInvRes listInvoicesResponse

		// decode/deserialize the response body into the struct we created
		if err = json.NewDecoder(response.Body).Decode(&lstInvRes); err != nil {
			log.Fatalf("could not decode response %v\n", err)
		}

		fmt.Printf("PS Status: %v. PS Message: %s \n", lstInvRes.Status, lstInvRes.Message)
		fmt.Printf("PS Meta: %v\n", lstInvRes.Meta)
		fmt.Printf("Number of invoices: %d\n", len(lstInvRes.Data))
	} else {
		fmt.Printf("Got response code %d from PS\n", response.StatusCode)
	}

}

func main() {
	downloadInvoices()
}
