package transaction

import (
	"time"

	"mpp/pkg/db"
	"mpp/pkg/model"
	validator "mpp/pkg/validate"

	"github.com/rs/xid"
)

func ACHTransaction(transaction model.Transaction) (model.Transaction, error) {

	transaction.ID = xid.New().String()
	transaction.PaymentMethodType = "ach"

	if err := validator.ValidateAch(transaction.PaymentMethod.Ach); err != nil {
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
