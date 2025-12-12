package ssl

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/sagar290/sslcommerz-go/helpers"
	"github.com/sagar290/sslcommerz-go/models"
)

const (
	SandboxURL = "https://sandbox.sslcommerz.com"
	LiveURL    = "https://securepay.sslcommerz.com"
)

type Ssl struct {
	StoreID                            string
	StorePassword                      string
	BaseUrl                            string
	InitURL                            string
	ValidationURL                      string
	RefundURL                          string
	RefundQueryURL                     string
	TransactionQueryBySessionIdURL     string
	TransactionQueryByTransactionIdURL string
}

type Option func(*Ssl)

// WithSandbox enables sandbox mode
func WithSandbox() Option {
	return func(s *Ssl) {
		s.BaseUrl = SandboxURL
		s.updateURLs()
	}
}

// WithBaseURL allows setting a custom base URL
func WithBaseURL(url string) Option {
	return func(s *Ssl) {
		s.BaseUrl = url
		s.updateURLs()
	}
}

func New(storeId string, storePassword string, opts ...Option) *Ssl {
	ssl := &Ssl{
		StoreID:       storeId,
		StorePassword: storePassword,
		BaseUrl:       LiveURL, // Default to Live
	}

	ssl.updateURLs()

	for _, opt := range opts {
		opt(ssl)
	}

	return ssl
}

func (s *Ssl) updateURLs() {
	s.InitURL = s.BaseUrl + "/gwprocess/v4/api.php?"
	s.ValidationURL = s.BaseUrl + "/validator/api/validationserverAPI.php?"
	s.RefundURL = s.BaseUrl + "/validator/api/merchantTransIDvalidationAPI.php?"
	s.RefundQueryURL = s.BaseUrl + "/validator/api/merchantTransIDvalidationAPI.php?"
	s.TransactionQueryBySessionIdURL = s.BaseUrl + "/validator/api/merchantTransIDvalidationAPI.php?"
	s.TransactionQueryByTransactionIdURL = s.BaseUrl + "/validator/api/merchantTransIDvalidationAPI.php?"
}

func (ssl *Ssl) MakePayment(req models.PaymentRequest) (models.PaymentResponse, error) {
	var res models.PaymentResponse

	// set store id and password
	req.StoreId = ssl.StoreID
	req.StorePasswd = ssl.StorePassword

	// Use go-querystring to encode the struct to url.Values
	values, err := query.Values(req)
	if err != nil {
		return models.PaymentResponse{}, err
	}

	response, err := http.PostForm(ssl.InitURL, values)
	if err != nil {
		log.Fatal(err)
		return models.PaymentResponse{}, err
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&res); err != nil {
		return models.PaymentResponse{}, err
	}

	return res, nil
}

func (ssl *Ssl) ValidatePayment(req models.OrderValidateRequest) (models.OrderValidateResponse, error) {
	var res models.OrderValidateResponse

	req.StoreId = ssl.StoreID
	req.StorePasswd = ssl.StorePassword

	v, _ := query.Values(req)

	resp, err := http.Get(ssl.ValidationURL + helpers.Encode(v))
	if err != nil {
		log.Fatal(err)
		return models.OrderValidateResponse{}, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return models.OrderValidateResponse{}, err
	}
	return res, nil
}

func (ssl *Ssl) InitiateRefund(req models.InitiateRefundRequest) (models.InitiateRefundResponse, error) {
	var res models.InitiateRefundResponse

	req.StoreId = ssl.StoreID
	req.StorePasswd = ssl.StorePassword

	v, _ := query.Values(req)

	resp, err := http.Get(ssl.RefundURL + helpers.Encode(v))
	if err != nil {
		log.Fatal(err)
		return models.InitiateRefundResponse{}, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return models.InitiateRefundResponse{}, err
	}
	return res, nil
}

func (ssl *Ssl) RefundQuery(req models.RefundQueryURLRequest) (models.RefundQueryURLResponse, error) {
	var res models.RefundQueryURLResponse

	req.StoreId = ssl.StoreID
	req.StorePasswd = ssl.StorePassword

	v, _ := query.Values(req)

	resp, err := http.Get(ssl.RefundQueryURL + helpers.Encode(v))
	if err != nil {
		log.Fatal(err)
		return models.RefundQueryURLResponse{}, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return models.RefundQueryURLResponse{}, err
	}
	return res, nil
}

func (ssl *Ssl) TransactionQueryBySessionId(req models.TransactionQueryBySessionKeyRequest) (models.TransactionQueryBySessionKeyResponse, error) {
	var res models.TransactionQueryBySessionKeyResponse

	req.StoreId = ssl.StoreID
	req.StorePasswd = ssl.StorePassword

	v, _ := query.Values(req)

	resp, err := http.Get(ssl.TransactionQueryBySessionIdURL + helpers.Encode(v))
	if err != nil {
		log.Fatal(err)
		return models.TransactionQueryBySessionKeyResponse{}, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return models.TransactionQueryBySessionKeyResponse{}, err
	}
	return res, nil
}

func (ssl *Ssl) TransactionQueryByTransactionId(req models.TransactionQueryByTransactionIdRequest) (models.TransactionQueryByTransactionIdResponse, error) {
	var res models.TransactionQueryByTransactionIdResponse

	req.StoreId = ssl.StoreID
	req.StorePasswd = ssl.StorePassword

	v, _ := query.Values(req)

	resp, err := http.Get(ssl.TransactionQueryBySessionIdURL + helpers.Encode(v))
	if err != nil {
		log.Fatal(err)
		return models.TransactionQueryByTransactionIdResponse{}, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return models.TransactionQueryByTransactionIdResponse{}, err
	}
	return res, nil
}

// ParseIPN parses the incoming IPN request and returns the IpnResponse struct
func ParseIPN(r *http.Request) (*models.IpnResponse, error) {
	if err := r.ParseForm(); err != nil {
		return nil, err
	}

	var ipn models.IpnResponse
	// Manually map form values to struct fields or use a decoder
	// Since we don't have a form decoder library, we'll do it manually for key fields
	// or use the same go-querystring library if it supports decoding (it usually doesn't).
	// For simplicity and robustness, let's map the most important fields manually.
	// Or better, use a loop with reflection if we want to be generic, but manual is safer for now.

	ipn.Status = r.FormValue("status")
	ipn.TranDate = r.FormValue("tran_date")
	ipn.TranId = r.FormValue("tran_id")
	ipn.ValId = r.FormValue("val_id")
	ipn.Amount = r.FormValue("amount")
	ipn.StoreAmount = r.FormValue("store_amount")
	ipn.CardType = r.FormValue("card_type")
	ipn.CardNo = r.FormValue("card_no")
	ipn.Currency = r.FormValue("currency")
	ipn.BankTranId = r.FormValue("bank_tran_id")
	ipn.CardIssuer = r.FormValue("card_issuer")
	ipn.CardBrand = r.FormValue("card_brand")
	ipn.CardIssuerCountry = r.FormValue("card_issuer_country")
	ipn.CardIssuerCountryCode = r.FormValue("card_issuer_country_code")
	ipn.CurrencyType = r.FormValue("currency_type")
	ipn.CurrencyAmount = r.FormValue("currency_amount")
	ipn.ValueA = r.FormValue("value_a")
	ipn.ValueB = r.FormValue("value_b")
	ipn.ValueC = r.FormValue("value_c")
	ipn.ValueD = r.FormValue("value_d")
	ipn.VerifySign = r.FormValue("verify_sign")
	ipn.VerifyKey = r.FormValue("verify_key")
	ipn.RiskLevel = r.FormValue("risk_level")
	ipn.RiskTitle = r.FormValue("risk_title")

	return &ipn, nil
}
