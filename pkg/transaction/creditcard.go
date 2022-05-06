package transaction

import (
	"time"

	"github.com/ngygbr/mpp/pkg/db"
	"github.com/ngygbr/mpp/pkg/model"
	validator "github.com/ngygbr/mpp/pkg/validate"

	"github.com/fluidpay/dough"
	"github.com/pkg/errors"
	"github.com/rs/xid"
)

func CCTransaction(transaction model.Transaction) (model.Transaction, error) {

	transaction.ID = xid.New().String()
	transaction.PaymentMethodType = "creditcard"

	if err := validator.ValidateCreditCard(transaction.PaymentMethod.CreditCard); err != nil {
		return model.Transaction{}, err
	}

	if err := validator.CheckIfSpecialCardNumber(transaction.PaymentMethod.CreditCard); err != nil {
		return model.Transaction{}, err
	}

	transaction.PaymentMethod.CreditCard.CardNumber = MaskCard(transaction.PaymentMethod.CreditCard)
	if transaction.PaymentMethod.CreditCard.CardNumber == "" {
		return model.Transaction{}, errors.New("can not mask card")
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

func MaskCard(card *model.CreditCard) string {
	_, _, maskedCard, err := dough.MaskCard(card.CardNumber)
	if err != nil {
		return ""
	}

	return maskedCard
}
