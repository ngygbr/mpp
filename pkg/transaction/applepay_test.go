package transaction

import (
	"github.com/stretchr/testify/assert"
	utils "github.com/ngygbr/mpp/pkg/config"
	"github.com/ngygbr/mpp/pkg/db"
	"github.com/ngygbr/mpp/pkg/model"
	"testing"
)

func TestApplePayTransaction(t *testing.T) {
	type args struct {
		transaction model.Transaction
	}
	tests := []struct {
		name string
		args args
	}{
		{"apple pay transaction test", args{transaction: model.Transaction{
			PaymentMethod: model.PaymentMethod{
				ApplePay: &model.ApplePay{
					PaymentToken: model.PaymentToken{
						Identifier:  "484d1cf96c8409e02c4c71276f265b65b8329bc1f8438cf66c08c975a7d4b84a",
						PaymentData: "38041f2368c5118806ed23951fe0f166e2f64099b6f6be495f5fbb248a154a0bf11e11a4bc47749d3e589eaeb59b428ae6b04ea1563140d5ef2118f623da8fdd06ed4c323560303d7ff1d15a5aacf6e93d9083fa21903ab5de65adbc3667a08cbe2cecb5beebbbe11cbdbebccad0d7e91d8f561f02466ffb70",
					},
				},
			},
			Amount: 1000,
			BillingAddress: model.Address{
				FirstName:    "Tester",
				LastName:     "Holder",
				PostalCode:   "5555",
				City:         "Szeged",
				AddressLine1: "Test street 1",
				Email:        "test@test.com",
				Phone:        "555555555",
			},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := utils.GetConfig()
			db.Connect(config.BadgerTmp + "/aptest")

			got, _ := ApplePayTransaction(tt.args.transaction)

			assert.NotZero(t, got.ID)
			assert.Equal(t, "pending_settlement", got.Status)
			assert.Equal(t, "apple_pay", got.PaymentMethodType)
			assert.NotZero(t, got.CreatedAt)
			assert.NotZero(t, got.UpdatedAt)

			assert.Equal(t, uint64(1000), got.Amount)
			assert.Equal(t, tt.args.transaction.BillingAddress, got.BillingAddress)
		})
	}
}
