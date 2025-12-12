package main

import (
	"fmt"
	"log"

	"github.com/sagar290/ssl_wireless_pgw_golang_sdk"
	"github.com/sagar290/ssl_wireless_pgw_golang_sdk/models"
)

func main() {
	// Initialize client with functional options
	// Default is Live, use WithSandbox() for sandbox mode
	ssl := ssl_wireless_pgw_golang_sdk.New("store_id", "store_password", ssl_wireless_pgw_golang_sdk.WithSandbox())

	// Create request using nested structs
	req := models.PaymentRequest{
		TranId: "123fgh",
		Shipping: models.Shipping{
			Method:   "online",
			Name:     "John Doe",
			Add1:     "Dhaka",
			City:     "Dhaka",
			State:    "Dhaka",
			Postcode: "1000",
			Country:  "Bangladesh",
		},
		Customer: models.Customer{
			Name:     "John Doe",
			Email:    "john@example.com",
			Add1:     "Dhaka",
			City:     "Dhaka",
			State:    "Dhaka",
			Postcode: "1000",
			Country:  "Bangladesh",
			Phone:    "01700000000",
		},
		ProductCategory: "Electronics",
		Product: models.Product{
			Name:    "Computer",
			Profile: "general",
		},
		SuccessUrl: "https://google.com",
		FailUrl:    "https://google.com",
		CancelUrl:  "https://google.com",
	}

	// Make payment
	response, err := ssl.MakePayment(req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Payment Response: %+v\n", response)
}
