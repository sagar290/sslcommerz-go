package main

import (
	"fmt"
	"log"

	ssl "github.com/sagar290/sslcommerz-go"
	"github.com/sagar290/sslcommerz-go/models"
)

func main() {
	// Initialize client with functional options
	// Default is Live, use WithSandbox() for sandbox mode
	client := ssl.New("store_id", "store_password", ssl.WithSandbox())

	// Create request using nested structs
	req := models.PaymentRequest{
		TranId: "123fgh",
		Shipping: models.Shipping{
			Method:   models.ShippingMethodOnline,
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
		ProductCategory: models.ProductCategoryElectronics,
		Product: models.Product{
			Name:    "Computer",
			Profile: models.ProductProfileGeneral,
		},
		SuccessUrl: "https://google.com",
		FailUrl:    "https://google.com",
		CancelUrl:  "https://google.com",
	}

	// Add Cart items
	err := req.SetCart([]models.CartProduct{
		{Product: "DHK TO BRS AC A1", Amount: "200.00"},
		{Product: "DHK TO BRS AC A2", Amount: "50.00"},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Make payment
	response, err := client.MakePayment(req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Payment Response: %+v\n", response)
}
