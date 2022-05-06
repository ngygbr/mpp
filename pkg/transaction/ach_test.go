package transaction

import (
	utils "github.com/ngygbr/mpp/pkg/config"
	"testing"

	"github.com/ngygbr/mpp/pkg/db"
	"github.com/ngygbr/mpp/pkg/model"

	"github.com/stretchr/testify/assert"
)

func TestACHTransaction(t *testing.T) {
	type args struct {
		transaction model.Transaction
	}
	tests := []struct {
		name    string
		args    args
	}{
		{"Ach transaction test", args{transaction: model.Transaction{
			PaymentMethod:     model.PaymentMethod{
				Ach: &model.Ach{
					AccountNumber: "123456789",
					RoutingNumber: "123456789",
					AccountType:   "checking",
					SECCode:       "web",
				},
			},
			Amount:            10000,
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
			db.Connect(config.BadgerTmp + "/achtest")

			got, _ := ACHTransaction(tt.args.transaction)

			assert.NotZero(t, got.ID)
			assert.Equal(t, "pending_settlement", got.Status)
			assert.Equal(t, "ach", got.PaymentMethodType)
			assert.NotZero(t, got.CreatedAt)
			assert.NotZero(t, got.UpdatedAt)

			assert.Equal(t, uint64(10000), got.Amount)
			assert.Equal(t, tt.args.transaction.BillingAddress, got.BillingAddress)
		})
	}
}
