package wrapper

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ssl "github.com/sagar290/sslcommerz-go"
	"github.com/sagar290/sslcommerz-go/models"
)

// IPNHandler returns a Gin handler that parses the IPN request and calls the provided callback.
func IPNHandler(callback func(*models.IpnResponse, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the IPN request using the SDK's helper
		ipn, err := ssl.ParseIPN(c.Request)
		if err != nil {
			// If parsing fails, return 400 Bad Request
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// Call the user's callback with the parsed IPN data
		callback(ipn, nil)

		// Acknowledge receipt (optional, but good practice)
		c.String(http.StatusOK, "IPN Received")
	}
}