package transaction

import (
	"time"

	"github.com/ngygbr/mpp/pkg/db"
	"github.com/ngygbr/mpp/pkg/model"
	validator "github.com/ngygbr/mpp/pkg/validate"

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
