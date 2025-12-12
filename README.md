# sslcommerz-go
> SSLCOMMERZ is the first payment gateway in Bangladesh opening doors for merchants to receive payments on the internet via their online stores. Their customers will be able to buy products online using their credit cards as well as bank accounts.

## Installation

```bash
go get github.com/sagar290/sslcommerz-go
```

## Usage

### Initialize Client

Use `New` to initialize the client. By default, it uses the Live environment.

```go
import (
    ssl "github.com/sagar290/sslcommerz-go"
    "github.com/sagar290/sslcommerz-go/models"
)

// Live Mode
client := ssl.New("store_id", "store_password")

// Sandbox Mode
client := ssl.New("store_id", "store_password", ssl.WithSandbox())
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
        Currency:        models.CurrencyBDT,
        ProductCategory: models.ProductCategoryElectronics,
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
            Method:   models.ShippingMethodCourier,
            NumItems: 1,
            Name:     "John Doe",
            Add1:     "Dhaka",
            City:     "Dhaka",
            Postcode: "1000",
            Country:  "Bangladesh",
        },
        
        Product: models.Product{
            Name:    "Computer",
            Profile: models.ProductProfileGeneral,
        },
    }

    // Add Cart items (Optional)
    err := req.SetCart([]models.CartProduct{
        {Product: "DHK TO BRS AC A1", Amount: "200.00"},
        {Product: "DHK TO BRS AC A2", Amount: "50.00"},
    })
    if err != nil {
        log.Fatal(err)
    }

    // Initiate Payment
    response, err := client.MakePayment(req)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Gateway Page URL: %s\n", response.GatewayPageURL)
}
```

### IPN Validation

You can parse the incoming IPN request using the `ParseIPN` helper.

```go
func IpnHandler(w http.ResponseWriter, r *http.Request) {
    // Parse IPN request
    ipn, err := ssl.ParseIPN(r)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Validate IPN (e.g., check status, amount, transaction ID)
    if ipn.Status == models.PaymentStatusValid {
        fmt.Printf("Payment Validated: %s\n", ipn.TranId)
    }
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