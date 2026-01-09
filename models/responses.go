package models

type PaymentResponse struct {
	Status       string `json:"status"`
	Failedreason string `json:"failedreason"`
	Sessionkey   string `json:"sessionkey"`
	Gw           struct {
		Visa            string `json:"visa"`
		Master          string `json:"master"`
		Amex            string `json:"amex"`
		Othercards      string `json:"othercards"`
		Internetbanking string `json:"internetbanking"`
		Mobilebanking   string `json:"mobilebanking"`
	} `json:"gw"`
	RedirectGatewayURL       string `json:"redirectGatewayURL"`
	DirectPaymentURLBank     string `json:"directPaymentURLBank"`
	DirectPaymentURLCard     string `json:"directPaymentURLCard"`
	DirectPaymentURL         string `json:"directPaymentURL"`
	RedirectGatewayURLFailed string `json:"redirectGatewayURLFailed"`
	GatewayPageURL           string `json:"GatewayPageURL"`
	StoreBanner              string `json:"storeBanner"`
	StoreLogo                string `json:"storeLogo"`
	Desc                     []struct {
		Name               string `json:"name"`
		Type               string `json:"type"`
		Logo               string `json:"logo"`
		Gw                 string `json:"gw"`
		RFlag              string `json:"r_flag,omitempty"`
		RedirectGatewayURL string `json:"redirectGatewayURL,omitempty"`
	} `json:"desc"`
	IsDirectPayEnable string `json:"is_direct_pay_enable"`
}

type IpnResponse struct {
	Status                string `json:"status"`
	TranDate              string `json:"tran_date"`
	TranId                string `json:"tran_id"`
	ValId                 string `json:"val_id"`
	Amount                string `json:"amount"`
	StoreAmount           string `json:"store_amount"`
	CardType              string `json:"card_type"`
	CardNo                string `json:"card_no"`
	Currency              string `json:"currency"`
	BankTranId            string `json:"bank_tran_id"`
	CardIssuer            string `json:"card_issuer"`
	CardBrand             string `json:"card_brand"`
	CardIssuerCountry     string `json:"card_issuer_country"`
	CardIssuerCountryCode string `json:"card_issuer_country_code"`
	CurrencyType          string `json:"currency_type"`
	CurrencyAmount        string `json:"currency_amount"`
	ValueA                string `json:"value_a"`
	ValueB                string `json:"value_b"`
	ValueC                string `json:"value_c"`
	ValueD                string `json:"value_d"`
	VerifySign            string `json:"verify_sign"`
	VerifyKey             string `json:"verify_key"`
	RiskLevel             string `json:"risk_level"`
	RiskTitle             string `json:"risk_title"`
}

// IsValid checks if the transaction status is valid
func (r *IpnResponse) IsValid() bool {
	return r.Status == PaymentStatusValid || r.Status == PaymentStatusValidated
}

type OrderValidateResponse struct {
	Status                string `json:"status"`
	TranDate              string `json:"tran_date"`
	TranId                string `json:"tran_id"`
	ValId                 string `json:"val_id"`
	Amount                string `json:"amount"`
	StoreAmount           string `json:"store_amount"`
	Currency              string `json:"currency"`
	BankTranId            string `json:"bank_tran_id"`
	CardType              string `json:"card_type"`
	CardNo                string `json:"card_no"`
	CardIssuer            string `json:"card_issuer"`
	CardBrand             string `json:"card_brand"`
	CardIssuerCountry     string `json:"card_issuer_country"`
	CardIssuerCountryCode string `json:"card_issuer_country_code"`
	CurrencyType          string `json:"currency_type"`
	CurrencyAmount        string `json:"currency_amount"`
	CurrencyRate          string `json:"currency_rate"`
	BaseFair              string `json:"base_fair"`
	ValueA                string `json:"value_a"`
	ValueB                string `json:"value_b"`
	ValueC                string `json:"value_c"`
	ValueD                string `json:"value_d"`
	EmiInstalment         string `json:"emi_instalment"`
	EmiAmount             string `json:"emi_amount"`
	EmiDescription        string `json:"emi_description"`
	EmiIssuer             string `json:"emi_issuer"`
	AccountDetails        string `json:"account_details"`
	RiskTitle             string `json:"risk_title"`
	RiskLevel             string `json:"risk_level"`
	APIConnect            string `json:"APIConnect"`
	ValidatedOn           string `json:"validated_on"`
	GwVersion             string `json:"gw_version"`
	DiscountPercentage    string `json:"discount_percentage,omitempty"`
	DiscountAmount        string `json:"discount_amount,omitempty"`
	DiscountRemarks       string `json:"discount_remarks,omitempty"`
	CartType              string `json:"cart_type,omitempty"`
	CartQuantity          string `json:"cart_quantity,omitempty"`
	AddressIsSame         string `json:"address_is_same,omitempty"`
}

// IsValid checks if the transaction status is valid
func (r *OrderValidateResponse) IsValid() bool {
	return r.Status == PaymentStatusValid || r.Status == PaymentStatusValidated
}

type TransactionQueryBySessionKeyResponse struct {
	Status                string `json:"status"`
	Sessionkey            string `json:"sessionkey"`
	TranDate              string `json:"tran_date"`
	TranId                string `json:"tran_id"`
	ValId                 string `json:"val_id"`
	Amount                string `json:"amount"`
	StoreAmount           string `json:"store_amount"`
	BankTranId            string `json:"bank_tran_id"`
	CardType              string `json:"card_type"`
	CardNo                string `json:"card_no"`
	CardIssuer            string `json:"card_issuer"`
	CardBrand             string `json:"card_brand"`
	CardIssuerCountry     string `json:"card_issuer_country"`
	CardIssuerCountryCode string `json:"card_issuer_country_code"`
	CurrencyType          string `json:"currency_type"`
	CurrencyAmount        string `json:"currency_amount"`
	CurrencyRate          string `json:"currency_rate"`
	BaseFair              string `json:"base_fair"`
	ValueA                string `json:"value_a"`
	ValueB                string `json:"value_b"`
	ValueC                string `json:"value_c"`
	ValueD                string `json:"value_d"`
	RiskTitle             string `json:"risk_title"`
	RiskLevel             string `json:"risk_level"`
	APIConnect            string `json:"APIConnect"`
	ValidatedOn           string `json:"validated_on"`
	GwVersion             string `json:"gw_version"`
}

type TransactionQueryByTransactionIdResponse struct {
	APIConnect     string `json:"APIConnect"`
	NoOfTransFound int    `json:"no_of_trans_found"`
	Element        []struct {
		ValId                 string `json:"val_id"`
		Status                string `json:"status"`
		ValidatedOn           string `json:"validated_on"`
		CurrencyType          string `json:"currency_type"`
		CurrencyAmount        string `json:"currency_amount"`
		CurrencyRate          string `json:"currency_rate"`
		BaseFair              string `json:"base_fair"`
		ValueA                string `json:"value_a"`
		ValueB                string `json:"value_b"`
		ValueC                string `json:"value_c"`
		ValueD                string `json:"value_d"`
		TranDate              string `json:"tran_date"`
		TranId                string `json:"tran_id"`
		Amount                string `json:"amount"`
		StoreAmount           string `json:"store_amount"`
		BankTranId            string `json:"bank_tran_id"`
		CardType              string `json:"card_type"`
		RiskTitle             string `json:"risk_title"`
		RiskLevel             string `json:"risk_level"`
		Currency              string `json:"currency"`
		BankGw                string `json:"bank_gw"`
		CardNo                string `json:"card_no"`
		CardIssuer            string `json:"card_issuer"`
		CardBrand             string `json:"card_brand"`
		CardIssuerCountry     string `json:"card_issuer_country"`
		CardIssuerCountryCode string `json:"card_issuer_country_code"`
		GwVersion             string `json:"gw_version"`
		EmiInstalment         string `json:"emi_instalment"`
		EmiAmount             string `json:"emi_amount"`
		EmiDescription        string `json:"emi_description"`
		EmiIssuer             string `json:"emi_issuer"`
		Error                 string `json:"error"`
	} `json:"element"`
}

type InitiateRefundResponse struct {
	APIConnect  string `json:"APIConnect"`
	BankTranId  string `json:"bank_tran_id"`
	TransId     string `json:"trans_id"`
	RefundRefId string `json:"refund_ref_id"`
	Status      string `json:"status"`
	ErrorReason string `json:"errorReason"`
}

type RefundQueryURLResponse struct {
	APIConnect  string `json:"APIConnect"`
	BankTranId  string `json:"bank_tran_id"`
	TranId      string `json:"tran_id"`
	InitiatedOn string `json:"initiated_on"`
	RefundedOn  string `json:"refunded_on"`
	Status      string `json:"status"`
	RefundRefId string `json:"refund_ref_id"`
}
