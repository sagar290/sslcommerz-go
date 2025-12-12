package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	ssl "github.com/sagar290/sslcommerz-go"
	"github.com/sagar290/sslcommerz-go/models"
	"github.com/sagar290/sslcommerz-go/wrapper"
)

func main() {
	r := gin.Default()

	// Initialize SSLCommerz client
	client := ssl.New("store_id", "store_password", ssl.WithSandbox())

	// Route to initiate payment
	r.GET("/pay", func(c *gin.Context) {
		req := models.PaymentRequest{
			TranId:          "123fgh",
			Currency:        models.CurrencyBDT,
			ProductCategory: models.ProductCategoryElectronics,
			SuccessUrl:      "http://localhost:8080/success",
			FailUrl:         "http://localhost:8080/fail",
			CancelUrl:       "http://localhost:8080/cancel",
			IpnUrl:          "http://localhost:8080/ipn",
			Customer: models.Customer{
				Name:  "John Doe",
				Email: "john@example.com",
			},
			Shipping: models.Shipping{
				Method: models.ShippingMethodOnline,
			},
			Product: models.Product{
				Name:    "Computer",
				Profile: models.ProductProfileGeneral,
			},
		}

		response, err := client.MakePayment(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Redirect user to GatewayPageURL
		c.Redirect(http.StatusFound, response.GatewayPageURL)
	})

	// IPN Handler Route using the wrapper
	r.POST("/ipn", wrapper.IPNHandler(func(ipn *models.IpnResponse, err error) {
		if err != nil {
			log.Printf("IPN Error: %v", err)
			return
		}
		if ipn.Status == models.PaymentStatusValid {
			log.Printf("Payment Validated: %s", ipn.TranId)
		}
	}))

	// Success Route
	r.POST("/success", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Payment Successful"})
	})

	// Fail Route
	r.POST("/fail", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Payment Failed"})
	})

	// Cancel Route
	r.POST("/cancel", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Payment Cancelled"})
	})

	log.Println("Server starting on :8080")
	r.Run(":8080")
}
