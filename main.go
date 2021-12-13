package main

import (
	"log"
	"os"

	"github.com/karosaxy/paystack-client/pkg/client/paystack"
)

const SECRET_KEY = "sk_test_3cec88e7ece6d9f5b69f33c013b94cc9142bf161"
const BASE_URL = "https://api.paystack.co"

func main() {

	secretKey := os.Getenv("PAYSTACK_SECRET_KEY")
	client := paystack.NewClient(BASE_URL, secretKey)

	ccq := paystack.CreateCustomerRequest{"man@go.com"}

	if err := client.CreateCustomer(ccq); err != nil {
		log.Fatalf("could not create customer. %v", err)
	}

	log.Println("Customer successfully created")

}
