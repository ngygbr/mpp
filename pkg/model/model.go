package model

import "time"

type Transaction struct {
	ID                string        `json:"id"`
	Status            string        `json:"status"`
	PaymentMethodType string        `json:"payment_method_type"`
	PaymentMethod     PaymentMethod `json:"payment_method"`
	Amount            uint64        `json:"amount"`
	BillingAddress    Address       `json:"billing_address"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at"`
}

type PaymentMethod struct {
	CreditCard *CreditCard `json:"credit_card,omitempty"`
	Ach        *Ach        `json:"ach,omitempty"`
	ApplePay   *ApplePay   `json:"apple_pay,omitempty"`
	GooglePay  *GooglePay  `json:"google_pay,omitempty"`
}

type Address struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	PostalCode   string `json:"postal_code"`
	City         string `json:"city"`
	AddressLine1 string `json:"address_line_1"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
}

type CreditCard struct {
	CardNumber     string  `json:"card_number"`
	HolderName     string  `json:"holder_name"`
	ExpirationDate *string `json:"exp_date,omitempty"`
	CVC            *string `json:"cvc,omitempty"`
}

type Ach struct {
	AccountNumber string  `json:"account_number"`
	RoutingNumber *string `json:"routing_number,omitempty"`
	AccountType   string  `json:"account_type,omitempty"`
	SECCode       *string `json:"sec_code"`
}

type ApplePay struct {
	PaymentToken PaymentToken `json:"payment_token"`
}

type GooglePay struct {
	EncryptedPayment EncryptedPayment `json:"encrypted_payment"`
}

type TestRequest struct {
	Method      string
	Path        string
	APIKey      string
	ContentType string
	Body        interface{}
}

type PaymentToken struct {
	Identifier  string `json:"identifier"`
	PaymentData string `json:"payment_data"`
}

type EncryptedPayment struct {
	PaymentID   string `json:"payment_id"`
	PaymentData string `json:"payment_data"`
}
