package models

import "encoding/json"

type CartProduct struct {
	Product string `json:"product"`
	Amount  string `json:"amount"`
}

// SetCart sets the cart parameter by marshaling a slice of CartProduct
func (p *PaymentRequest) SetCart(products []CartProduct) error {
	data, err := json.Marshal(products)
	if err != nil {
		return err
	}
	p.Cart = string(data)
	return nil
}
