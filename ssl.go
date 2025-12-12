package ssl_wireless_pgw_golang_sdk

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/sagar290/ssl_wireless_pgw_golang_sdk/helpers"
	"github.com/sagar290/ssl_wireless_pgw_golang_sdk/models"
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
