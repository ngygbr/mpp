package transaction

import (
	"encoding/json"
	"errors"
	"log"
	crypter "github.com/ngygbr/mpp/pkg/crypto"
	"time"

	"github.com/ngygbr/mpp/pkg/db"
	"github.com/ngygbr/mpp/pkg/model"
	validator "github.com/ngygbr/mpp/pkg/validate"

	"github.com/rs/xid"
)

func GooglePayTransaction(transaction model.Transaction) (model.Transaction, error) {

	transaction.ID = xid.New().String()
	transaction.PaymentMethodType = "google_pay"

	if err := validator.ValidateGooglePay(transaction.PaymentMethod.GooglePay); err != nil {
		return model.Transaction{}, err
	}

	cipherText := transaction.PaymentMethod.GooglePay.EncryptedPayment.PaymentData
	decryptKey := transaction.PaymentMethod.GooglePay.EncryptedPayment.PaymentID

	cardDataStringFormat, err := crypter.DecryptCard(cipherText, decryptKey)
	if err != nil {
		return model.Transaction{}, err
	}

	var cardData model.CreditCard
	err = json.Unmarshal([]byte(cardDataStringFormat), &cardData)
	if err != nil {
		log.Fatalf("Error occured during unmarshaling. Error: %s", err.Error())
	}

	if err := validator.ValidateCreditCard(&cardData); err != nil {
		return model.Transaction{}, err
	}

	cardData.CardNumber = MaskCard(&cardData)
	if cardData.CardNumber == "" {
		return model.Transaction{}, errors.New("can not mask card")
	}

	transaction.PaymentMethod = model.PaymentMethod{
		CreditCard: &cardData,
	}

	transaction.Status = "pending_settlement"
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()

	err = db.Create(&transaction)
	if err != nil {
		return model.Transaction{}, err
	}

	return transaction, nil
}
