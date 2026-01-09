package models

// PaymentRequest represents the body of the payment request
type PaymentRequest struct {
	// Basic
	StoreId         string  `json:"store_id" url:"store_id"`
	StorePasswd     string  `json:"store_passwd" url:"store_passwd"`
	TotalAmount     float64 `json:"total_amount" url:"total_amount"`
	Currency        string  `json:"currency" url:"currency"` // default:"BDT"
	TranId          string  `json:"tran_id" url:"tran_id"`
	ProductCategory string  `json:"product_category" url:"product_category"`
	SuccessUrl      string  `json:"success_url" url:"success_url"`
	FailUrl         string  `json:"fail_url" url:"fail_url"`
	CancelUrl       string  `json:"cancel_url" url:"cancel_url"`
	IpnUrl          string  `json:"ipn_url" url:"ipn_url"`

	// Card Details
	MultiCardName    string `json:"multi_card_name" url:"multi_card_name,omitempty"`
	AllowedBin       string `json:"allowed_bin" url:"allowed_bin,omitempty"`
	EmiOption        string `json:"emi_option" url:"emi_option,omitempty"`
	EmiMaxInstOption string `json:"emi_max_inst_option" url:"emi_max_inst_option,omitempty"`
	EmiSelectedInst  string `json:"emi_selected_inst" url:"emi_selected_inst,omitempty"`
	EmiAllowOnly     string `json:"emi_allow_only" url:"emi_allow_only,omitempty"`

	// Nested Structs
	Customer `url:",squash"`
	Shipping `url:",squash"`
	Product  `url:",squash"`
	Extra    `url:",squash"`
}

type Customer struct {
	Name     string `json:"cus_name" url:"cus_name"`
	Email    string `json:"cus_email" url:"cus_email"`
	Add1     string `json:"cus_add1" url:"cus_add1"`
	Add2     string `json:"cus_add2" url:"cus_add2"`
	City     string `json:"cus_city" url:"cus_city"`
	State    string `json:"cus_state" url:"cus_state"`
	Postcode string `json:"cus_postcode" url:"cus_postcode"`
	Country  string `json:"cus_country" url:"cus_country"`
	Phone    string `json:"cus_phone" url:"cus_phone"`
	Fax      string `json:"cus_fax" url:"cus_fax,omitempty"`
}

type Shipping struct {
	Method   string `json:"shipping_method" url:"shipping_method"`
	NumItems int    `json:"num_of_item" url:"num_of_item"`
	Name     string `json:"ship_name" url:"ship_name"`
	Add1     string `json:"ship_add1" url:"ship_add1"`
	Add2     string `json:"ship_add2" url:"ship_add2"`
	City     string `json:"ship_city" url:"ship_city"`
	State    string `json:"ship_state" url:"ship_state"`
	Postcode string `json:"ship_postcode" url:"ship_postcode"`
	Country  string `json:"ship_country" url:"ship_country"`
}

type Product struct {
	Name               string `json:"product_name" url:"product_name"`
	Profile            string `json:"product_profile" url:"product_profile"`
	HoursTillDeparture string `json:"hours_till_departure" url:"hours_till_departure,omitempty"`
	FlightType         string `json:"flight_type" url:"flight_type,omitempty"`
	Pnr                string `json:"pnr" url:"pnr,omitempty"`
	JourneyFromTo      string `json:"journey_from_to" url:"journey_from_to,omitempty"`
	ThirdPartyBooking  string `json:"third_party_booking" url:"third_party_booking,omitempty"`
	HotelName          string `json:"hotel_name" url:"hotel_name,omitempty"`
	LengthOfStay       string `json:"length_of_stay" url:"length_of_stay,omitempty"`
	CheckInTime        string `json:"check_in_time" url:"check_in_time,omitempty"`
	HotelCity          string `json:"hotel_city" url:"hotel_city,omitempty"`
	Type               string `json:"product_type" url:"product_type,omitempty"`
	TopupNumber        string `json:"topup_number" url:"topup_number,omitempty"`
	CountryTopup       string `json:"country_topup" url:"country_topup,omitempty"`
	Cart               string `json:"cart" url:"cart,omitempty"`
	Amount             string `json:"product_amount" url:"product_amount,omitempty"`
	Vat                string `json:"vat" url:"vat,omitempty"`
	DiscountAmount     string `json:"discount_amount" url:"discount_amount,omitempty"`
	ConvenienceFee     string `json:"convenience_fee" url:"convenience_fee,omitempty"`
}

type Extra struct {
	ValueA string `json:"value_a" url:"value_a,omitempty"`
	ValueB string `json:"value_b" url:"value_b,omitempty"`
	ValueC string `json:"value_c" url:"value_c,omitempty"`
	ValueD string `json:"value_d" url:"value_d,omitempty"`
}

// Constants for PaymentRequest
const (
	// Shipping Methods
	ShippingMethodOnline  = "online"
	ShippingMethodCourier = "courier"
	ShippingMethodNo      = "no"

	// Product Categories
	ProductCategoryElectronics = "Electronics"
	ProductCategoryClothing    = "Clothing"
	ProductCategoryDigital     = "Digital"

	// Product Profiles
	ProductProfileGeneral          = "general"
	ProductProfilePhysicalGoods    = "physical-goods"
	ProductProfileNonPhysicalGoods = "non-physical-goods"
	ProductProfileAirlineTickets   = "airline-tickets"
	ProductProfileTravelVertical   = "travel-vertical"
	ProductProfileTelecomVertical  = "telecom-vertical"

	// Payment Status
	PaymentStatusValid              = "VALID"
	PaymentStatusValidated          = "VALIDATED"
	PaymentStatusInvalidTransaction = "INVALID_TRANSACTION"

	PaymentStatusFailed      = "FAILED"
	PaymentStatusCancelled   = "CANCELLED"
	PaymentStatusUnattempted = "UNATTEMPTED"
	PaymentStatusExpired     = "EXPIRED"

	// Currencies
	CurrencyBDT = "BDT"
	CurrencyUSD = "USD"
	CurrencyEUR = "EUR"
	CurrencyGBP = "GBP"
)
