package validpackage

import (
	"testing"

	"github.com/ngygbr/mpp/pkg/model"

	"github.com/pkg/errors"
)

func TestValidateCreditCard(t *testing.T) {
	tests := []struct {
		name    string
		card *model.CreditCard
		wantErr bool
	}{
		{"valid", &model.CreditCard{
			CardNumber:     "4111111111111111",
			HolderName:     "test holder name",
			ExpirationDate: "05/25",
			CVC:            "111",
		}, false},
		{"invalid luhn", &model.CreditCard{
			CardNumber:     "6969696969696969",
			HolderName:     "test holder name",
			ExpirationDate: "05/25",
			CVC:            "111",
		}, true},
		{"invalid card number regex", &model.CreditCard{
			CardNumber:     "1",
			HolderName:     "test holder name",
			ExpirationDate: "05/25",
			CVC:            "111",
		}, true},
		{"invalid holder name", &model.CreditCard{
			CardNumber:     "4111111111111111",
			HolderName:     "1",
			ExpirationDate: "05/25",
			CVC:            "111",
		}, true},
		{"invalid expiration date regex", &model.CreditCard{
			CardNumber:     "4111111111111111",
			HolderName:     "test holder name",
			ExpirationDate: "",
			CVC:            "111",
		}, true},
		{"expired expiration date", &model.CreditCard{
			CardNumber:     "4111111111111111",
			HolderName:     "test holder name",
			ExpirationDate: "05/20",
			CVC:            "111",
		}, true},
		{"invalid cvc", &model.CreditCard{
			CardNumber:     "4111111111111111",
			HolderName:     "test holder name",
			ExpirationDate: "05/25",
			CVC:            "",
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateCreditCard(tt.card); (err != nil) != tt.wantErr {
				t.Errorf("ValidateCreditCard() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckIfSpecialCardNumber(t *testing.T) {
	tests := []struct {
		name    string
		card *model.CreditCard
		errMsg error
	}{
		{"limit exceeded", &model.CreditCard{
			CardNumber:     "4455444455551111",
			HolderName:     "test holder name",
			ExpirationDate: "05/25",
			CVC:            "111",
		}, errors.New("limit exceeded")},
		{"daily limit exceeded", &model.CreditCard{
			CardNumber:     "7755444455551111",
			HolderName:     "test holder name",
			ExpirationDate: "05/25",
			CVC:            "111",
		}, errors.New("daily limit exceeded")},
		{"fraud detected", &model.CreditCard{
			CardNumber:     "8888888888888888",
			HolderName:     "test holder name",
			ExpirationDate: "05/25",
			CVC:            "111",
		}, errors.New("fraud detected")},
		{"card blocked", &model.CreditCard{
			CardNumber:     "0000000000000000",
			HolderName:     "test holder name",
			ExpirationDate: "05/25",
			CVC:            "111",
		}, errors.New("card blocked")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckIfSpecialCardNumber(tt.card); err.Error() != tt.errMsg.Error() {
				t.Errorf("got error should be %v, but expected error msg was %v", err, tt.errMsg)
			}
		})
	}
}

func TestValidateAch(t *testing.T) {
	tests := []struct {
		name    string
		ach *model.Ach
		wantErr bool
	}{
		{"valid", &model.Ach{
			AccountNumber: "111111111111111",
			RoutingNumber: "111111111",
			AccountType:   "checking",
			SECCode:       "web",
		}, false},
		{"invalid account number", &model.Ach{
			AccountNumber: "invalid",
			RoutingNumber: "111111111",
			AccountType:   "checking",
			SECCode:       "web",
		}, true},
		{"invalid routing number", &model.Ach{
			AccountNumber: "111111111111111",
			RoutingNumber: "invalid",
			AccountType:   "checking",
			SECCode:       "web",
		}, true},
		{"invalid account type", &model.Ach{
			AccountNumber: "111111111111111",
			RoutingNumber: "111111111",
			AccountType:   "invalid",
			SECCode:       "web",
		}, true},
		{"invalid sec code", &model.Ach{
			AccountNumber: "111111111111111",
			RoutingNumber: "111111111",
			AccountType:   "checking",
			SECCode:       "invalid",
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateAch(tt.ach); (err != nil) != tt.wantErr {
				t.Errorf("ValidateAch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateAddress(t *testing.T) {
	tests := []struct {
		name    string
		address *model.Address
		wantErr bool
	}{
		{"valid address", &model.Address{
			FirstName:    "Test",
			LastName:     "Name",
			PostalCode:   "6000",
			City:         "Test City",
			AddressLine1: "Test street 1",
			Email:        "test@test.com",
			Phone:        "555555555",
		}, false},
		{"invalid first name", &model.Address{
			FirstName:    "1",
			LastName:     "Name",
			PostalCode:   "6000",
			City:         "Test City",
			AddressLine1: "Test street 1",
			Email:        "test@test.com",
			Phone:        "555555555",
		}, true},
		{"invalid last name", &model.Address{
			FirstName:    "Test",
			LastName:     "1",
			PostalCode:   "6000",
			City:         "Test City",
			AddressLine1: "Test street 1",
			Email:        "test@test.com",
			Phone:        "555555555",
		}, true},
		{"invalid postal code", &model.Address{
			FirstName:    "Test",
			LastName:     "Name",
			PostalCode:   "1",
			City:         "Test City",
			AddressLine1: "Test street 1",
			Email:        "test@test.com",
			Phone:        "555555555",
		}, true},
		{"invalid city", &model.Address{
			FirstName:    "Test",
			LastName:     "Name",
			PostalCode:   "6000",
			City:         "1",
			AddressLine1: "Test street 1",
			Email:        "test@test.com",
			Phone:        "555555555",
		}, true},
		{"invalid address line 1", &model.Address{
			FirstName:    "Test",
			LastName:     "Name",
			PostalCode:   "6000",
			City:         "Test City",
			AddressLine1: "1",
			Email:        "test@test.com",
			Phone:        "555555555",
		}, true},
		{"invalid email", &model.Address{
			FirstName:    "Test",
			LastName:     "Name",
			PostalCode:   "6000",
			City:         "Test City",
			AddressLine1: "Test street 1",
			Email:        "1",
			Phone:        "555555555",
		}, true},
		{"invalid phone", &model.Address{
			FirstName:    "Test",
			LastName:     "Name",
			PostalCode:   "6000",
			City:         "Test City",
			AddressLine1: "Test street 1",
			Email:        "test@test.com",
			Phone:        "0",
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateAddress(tt.address); (err != nil) != tt.wantErr {
				t.Errorf("ValidateAddress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateXID(t *testing.T) {
	tests := []struct {
		name    string
		xid string
		wantErr bool
	}{
		{"valid xid", "aaaaaaaaaaaaaaaaaaaa", false},
		{"invalid xid", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateXID(tt.xid); (err != nil) != tt.wantErr {
				t.Errorf("ValidateXID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidatePaymentMethodType(t *testing.T) {
	tests := []struct {
		name    string
		paymentMethodType string
		wantErr bool
	}{
		{"valid card payment method type", "creditcard", false},
		{"valid ach payment method type", "ach", false},
		{"valid apple_pay payment method type", "apple_pay", false},
		{"valid google_pay payment method type", "google_pay", false},
		{"valid apm payment method type", "apm", false},
		{"invalid payment method type", "invalid", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidatePaymentMethodType(tt.paymentMethodType); (err != nil) != tt.wantErr {
				t.Errorf("ValidatePaymentMethodType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidatePaymentMethod(t *testing.T) {
	tests := []struct {
		name    string
		paymentMethod *model.PaymentMethod
		wantErr bool
	}{
		{"valid cc payment method", &model.PaymentMethod{
			CreditCard: &model.CreditCard{
				CardNumber:     "4111111111111111",
				HolderName:     "Test Holder",
				ExpirationDate: "05/25",
				CVC:            "444",
			},
		},false},
		{"valid ach payment method", &model.PaymentMethod{
			Ach: &model.Ach{
				AccountNumber: "123456789",
				RoutingNumber: "123456789",
				AccountType:   "checking",
				SECCode:       "web",
			},
		},false},
		{"valid apple pay payment method", &model.PaymentMethod{
			ApplePay: &model.ApplePay{
				PaymentToken: model.PaymentToken{
					Identifier:  "484d1cf96c8409e02c4c71276f265b65b8329bc1f8438cf66c08c975a7d4b84a",
					PaymentData: "38041f2368c5118806ed23951fe0f166e2f64099b6f6be495f5fbb248a154a0bf11e11a4bc47749d3e589eaeb59b428ae6b04ea1563140d5ef2118f623da8fdd06ed4c323560303d7ff1d15a5aacf6e93d9083fa21903ab5de65adbc3667a08cbe2cecb5beebbbe11cbdbebccad0d7e91d8f561f02466ffb70",
				},
			},
		},false},
		{"valid google pay payment method", &model.PaymentMethod{
			GooglePay: &model.GooglePay{
				EncryptedPayment: model.EncryptedPayment{
					PaymentID:   "484d1cf96c8409e02c4c71276f265b65b8329bc1f8438cf66c08c975a7d4b84a",
					PaymentData: "38041f2368c5118806ed23951fe0f166e2f64099b6f6be495f5fbb248a154a0bf11e11a4bc47749d3e589eaeb59b428ae6b04ea1563140d5ef2118f623da8fdd06ed4c323560303d7ff1d15a5aacf6e93d9083fa21903ab5de65adbc3667a08cbe2cecb5beebbbe11cbdbebccad0d7e91d8f561f02466ffb70",
				},
			},
		},false},
		{"invalid cc and ach payment method", &model.PaymentMethod{
			CreditCard: &model.CreditCard{
				CardNumber:     "4111111111111111",
				HolderName:     "Test Holder",
				ExpirationDate: "05/25",
				CVC:            "444",
			},
			Ach: &model.Ach{
				AccountNumber: "123456789",
				RoutingNumber: "123456789",
				AccountType:   "checking",
				SECCode:       "web",
			},
		},true},
		{"invalid cc and apple pay payment method", &model.PaymentMethod{
			CreditCard: &model.CreditCard{
				CardNumber:     "4111111111111111",
				HolderName:     "Test Holder",
				ExpirationDate: "05/25",
				CVC:            "444",
			},
			ApplePay: &model.ApplePay{
				PaymentToken: model.PaymentToken{
					Identifier:  "",
					PaymentData: "",
				},
			},
		},true},
		{"invalid ach and google pay payment method", &model.PaymentMethod{
			Ach: &model.Ach{
				AccountNumber: "123456789",
				RoutingNumber: "123456789",
				AccountType:   "checking",
				SECCode:       "web",
			},
			GooglePay: &model.GooglePay{
				EncryptedPayment: model.EncryptedPayment{
					PaymentID:   "484d1cf96c8409e02c4c71276f265b65b8329bc1f8438cf66c08c975a7d4b84a",
					PaymentData: "",
				},
			},
		},true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidatePaymentMethod(tt.paymentMethod); (err != nil) != tt.wantErr {
				t.Errorf("ValidatePaymentMethod() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateApplePay(t *testing.T) {
	tests := []struct {
		name    string
		ap *model.ApplePay
		wantErr bool
	}{
		{"Valid apple pay", &model.ApplePay{
			PaymentToken: model.PaymentToken{
				Identifier:  "484d1cf96c8409e02c4c71276f265b65b8329bc1f8438cf66c08c975a7d4b84a",
				PaymentData: "38041f2368c5118806ed23951fe0f166e2f64099b6f6be495f5fbb248a154a0bf11e11a4bc47749d3e589eaeb59b428ae6b04ea1563140d5ef2118f623da8fdd06ed4c323560303d7ff1d15a5aacf6e93d9083fa21903ab5de65adbc3667a08cbe2cecb5beebbbe11cbdbebccad0d7e91d8f561f02466ffb70",
			},
		}, false},
		{"Invalid apple pay identifier", &model.ApplePay{
			PaymentToken: model.PaymentToken{
				Identifier:  "",
				PaymentData: "38041f2368c5118806ed23951fe0f166e2f64099b6f6be495f5fbb248a154a0bf11e11a4bc47749d3e589eaeb59b428ae6b04ea1563140d5ef2118f623da8fdd06ed4c323560303d7ff1d15a5aacf6e93d9083fa21903ab5de65adbc3667a08cbe2cecb5beebbbe11cbdbebccad0d7e91d8f561f02466ffb70",
			},
		}, false},
		{"Inalid apple pay paymentData", &model.ApplePay{
			PaymentToken: model.PaymentToken{
				Identifier:  "484d1cf96c8409e02c4c71276f265b65b8329bc1f8438cf66c08c975a7d4b84a",
				PaymentData: "",
			},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateApplePay(tt.ap); (err != nil) != tt.wantErr {
				t.Errorf("ValidateApplePay() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateGooglePay(t *testing.T) {
	tests := []struct {
		name    string
		gp *model.GooglePay
		wantErr bool
	}{
		{"Valid google pay", &model.GooglePay{
			EncryptedPayment: model.EncryptedPayment{
				PaymentID:   "484d1cf96c8409e02c4c71276f265b65b8329bc1f8438cf66c08c975a7d4b84a",
				PaymentData: "38041f2368c5118806ed23951fe0f166e2f64099b6f6be495f5fbb248a154a0bf11e11a4bc47749d3e589eaeb59b428ae6b04ea1563140d5ef2118f623da8fdd06ed4c323560303d7ff1d15a5aacf6e93d9083fa21903ab5de65adbc3667a08cbe2cecb5beebbbe11cbdbebccad0d7e91d8f561f02466ffb70",
			},
		}, false},
		{"Invalid google pay payment id", &model.GooglePay{
			EncryptedPayment: model.EncryptedPayment{
				PaymentID:   "",
				PaymentData: "38041f2368c5118806ed23951fe0f166e2f64099b6f6be495f5fbb248a154a0bf11e11a4bc47749d3e589eaeb59b428ae6b04ea1563140d5ef2118f623da8fdd06ed4c323560303d7ff1d15a5aacf6e93d9083fa21903ab5de65adbc3667a08cbe2cecb5beebbbe11cbdbebccad0d7e91d8f561f02466ffb70",
			},
		}, false},
		{"Invalid google pay payment data", &model.GooglePay{
			EncryptedPayment: model.EncryptedPayment{
				PaymentID:   "484d1cf96c8409e02c4c71276f265b65b8329bc1f8438cf66c08c975a7d4b84a",
				PaymentData: "",
			},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateGooglePay(tt.gp); (err != nil) != tt.wantErr {
				t.Errorf("ValidateGooglePay() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
