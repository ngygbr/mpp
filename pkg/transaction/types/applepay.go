package types

import (
	"github.com/rs/xid"
	"mpp/pkg/db"
	"mpp/pkg/model"
	validate "mpp/pkg/validate"
	"time"
)

func ApplePayTransaction(transaction model.Transaction) (model.Transaction, error){

	transaction.ID = xid.New().String()
	transaction.PaymentMethodType = "apple_pay"

	if err := validate.ValidateApplePay(transaction.PaymentMethod.ApplePay); err != nil {
		return model.Transaction{}, err
	}

	if err := validate.ValidateAddress(&transaction.BillingAddress); err != nil {
		return model.Transaction{}, err
	}

	transaction.Status = "pending_settlement"
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()

	err := db.Create(&transaction)
	if err != nil {
		return model.Transaction{}, err
	}

	return transaction, nil
}
