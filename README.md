# ssl-wireless-pgw-golang-sdk (Beta)
> SSLCOMMERZ is the first payment gateway in Bangladesh opening doors for merchants to receive payments on the internet via their online stores. Their customers will be able to buy products online using their credit cards as well as bank accounts.

## Installation

```bash
go get github.com/sagar290/ssl_wireless_pgw_golang_sdk
```

## Usage

### Initialize Client

Use `New` to initialize the client. By default, it uses the Live environment.

```go
import (
    "github.com/sagar290/ssl_wireless_pgw_golang_sdk"
    "github.com/sagar290/ssl_wireless_pgw_golang_sdk/models"
)

// Live Mode
ssl := ssl_wireless_pgw_golang_sdk.New("store_id", "store_password")

// Sandbox Mode
ssl := ssl_wireless_pgw_golang_sdk.New("store_id", "store_password", ssl_wireless_pgw_golang_sdk.WithSandbox())
```

### Make Payment

Construct the `PaymentRequest` using nested structs for a clean and readable initialization.

```go
func main() {
    // ... initialize ssl client ...

    // Create request
    req := models.PaymentRequest{
        TranId:          "123fgh",
        TotalAmount:     250,
        Currency:        "BDT",
        ProductCategory: "Electronics",
        SuccessUrl:      "https://example.com/success",
        FailUrl:         "https://example.com/fail",
        CancelUrl:       "https://example.com/cancel",
        IpnUrl:          "https://example.com/ipn",
        
        Customer: models.Customer{
            Name:     "John Doe",
            Email:    "john@example.com",
            Phone:    "01700000000",
            Add1:     "Dhaka",
            City:     "Dhaka",
            Postcode: "1000",
            Country:  "Bangladesh",
        },
        
        Shipping: models.Shipping{
            Method:   "Courier",
            NumItems: 1,
            Name:     "John Doe",
            Add1:     "Dhaka",
            City:     "Dhaka",
            Postcode: "1000",
            Country:  "Bangladesh",
        },
        
        Product: models.Product{
            Name:    "Computer",
            Profile: "general",
        },
    }

    // Initiate Payment
    response, err := ssl.MakePayment(req)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Gateway Page URL: %s\n", response.GatewayPageURL)
}
```

### Other Methods

The SDK supports other operations like validation, refund, and transaction queries.

- `ValidatePayment(req models.OrderValidateRequest)`
- `InitiateRefund(req models.InitiateRefundRequest)`
- `RefundQuery(req models.RefundQueryURLRequest)`
- `TransactionQueryBySessionId(req models.TransactionQueryBySessionKeyRequest)`
- `TransactionQueryByTransactionId(req models.TransactionQueryByTransactionIdRequest)`

## Contribution guideline
Make these things more useful by contributing our brain and efforts. History will remember our contribution as far as lazy peoples like me are alive ;)

- Clone the repository
- Create branch from master
- make changes
- push the branch