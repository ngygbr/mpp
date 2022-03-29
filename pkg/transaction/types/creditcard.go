package types

import (
	"github.com/pkg/errors"
	"mock-paymentprocessor/pkg/db"
	"mock-paymentprocessor/pkg/model"
	validate "mock-paymentprocessor/pkg/validate"
	"time"

	"github.com/fluidpay/dough"
	"github.com/rs/xid"
)

func CCTransaction(transaction model.Transaction) (model.Transaction, error) {

	transaction.ID = xid.New().String()
	transaction.PaymentMethodType = "creditcard"

	if err := validate.SpecialCardNumbers(transaction.PaymentMethod.CreditCard); err != nil {
		return model.Transaction{}, err
	}

	if err := validate.ValidateCreditCard(transaction.PaymentMethod.CreditCard); err != nil {
		return model.Transaction{}, err
	}

	if err := validate.ValidateAddress(&transaction.BillingAddress); err != nil {
		return model.Transaction{}, err
	}

	transaction.PaymentMethod.CreditCard.CardNumber = maskCard(transaction.PaymentMethod.CreditCard)
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

func maskCard(card *model.CreditCard) string {
	_, _, maskedCard, err := dough.MaskCard(card.CardNumber)
	if err != nil {
		return ""
	}

	return maskedCard
}
