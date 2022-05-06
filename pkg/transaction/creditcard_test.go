package transaction

import (
	"github.com/stretchr/testify/assert"
	utils "github.com/ngygbr/mpp/pkg/config"
	"github.com/ngygbr/mpp/pkg/db"
	"github.com/ngygbr/mpp/pkg/model"
	"testing"
)

func TestCCTransaction(t *testing.T) {
	type args struct {
		transaction model.Transaction
	}
	tests := []struct {
		name    string
		args    args
	}{
		{"Credit card transaction test", args{transaction: model.Transaction{
			PaymentMethod:     model.PaymentMethod{
				CreditCard: &model.CreditCard{
					CardNumber:     "4111111111111111",
					HolderName:     "Tester Holder",
					ExpirationDate: "06/25",
					CVC:            "444",
				},
			},
			Amount:            1000,
			BillingAddress:    model.Address{
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
			db.Connect(config.BadgerTmp + "/cctest")

			got, _ := CCTransaction(tt.args.transaction)

			assert.NotZero(t, got.ID)
			assert.Equal(t, "pending_settlement", got.Status)
			assert.Equal(t, "creditcard", got.PaymentMethodType)
			assert.NotZero(t, got.CreatedAt)
			assert.NotZero(t, got.UpdatedAt)

			assert.Equal(t, uint64(1000), got.Amount)
			assert.Equal(t, tt.args.transaction.BillingAddress, got.BillingAddress)
		})
	}
}
