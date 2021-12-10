package paystack

import "errors"


var ErrCustomerNotFound = errors.New("Customer with ID/Email not found")

type Customer struct {
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

type Amount int

type DateTime string

type LineItem struct {
	Name   string `json:"name"`
	Amount Amount `json:"amount"`
}

type Tax struct {
	Name   string `json:"name"`
	Amount Amount `json:"amount"`
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
	LineItems     []LineItem `json:"line_items"`
	Tax           []Tax      `json:"tax"`
	Customer      Customer   `json:"customer"`
	RequestCode   string     `json:"request_code"`
	Status        string     `json:"status"`
	Paid          bool       `json:"paid"`
	PaidAt        DateTime   `json:"paid_at"`
	CreatedAt     DateTime   `json:"created_at"`
}
type Meta struct {
	Total     int `json:"total"`
	Skipped   int `json:"skipped"`
	PerPage   int `json:"perPage"`
	Page      int `json:"page"`
	PageCount int `json:"pageCount"`
}

type ListInvoicesResponse struct {
	Status  bool      `json:"status"`
	Message string    `json:"message"`
	Meta    Meta      `json:"meta"`
	Data    []Invoice `json:"data"`
}

type GetCustomerResponse struct {
	Customer Customer `json:"data"`
}

type CreateCustomerRequest struct {
	Email string `json:"email"`
}