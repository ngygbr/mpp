package validpackage

import (
	configuration "mpp/pkg/config"
	"strconv"
	"strings"
	"time"

	"mpp/pkg/model"

	"github.com/fluidpay/dough"
	"github.com/pkg/errors"
)

var config = configuration.GetConfig()

func ValidateCreditCard(card *model.CreditCard) error {

	if !dough.ValidLuhn(card.CardNumber) {
		return errors.New("invalid card number for luhn")
	}

	if !cardNumberRegex.Match([]byte(card.CardNumber)) {
		return errors.New("invalid card number for regex")
	}

	if !cvcRegex.Match([]byte(card.CVC)) {
		return errors.New("invalid cvc code")
	}

	if !holderNameRegex.Match([]byte(card.HolderName)) {
		return errors.New("invalid holder name")
	}

	if !expDateRegex.Match([]byte(card.ExpirationDate)) {
		return errors.New("invalid expiration date for regex")
	}

	month, _ := strconv.Atoi(strings.Split(card.ExpirationDate, "/")[0])
	year, _ := strconv.Atoi("20" + strings.Split(card.ExpirationDate, "/")[1])
	expDate := time.Date(year, time.Month(month), 0, 0, 0, 0, 0, time.UTC)

	if !expDate.After(time.Now()) {
		return errors.New("expired expiration date")
	}

	return nil
}

func CheckIfSpecialCardNumber(card *model.CreditCard) error {

	if !config.DisableLimit {
		if card.CardNumber == "4455444455551111" {
			return errors.New("limit exceeded")
		}
	}

	if !config.DisableDailyLimit {
		if card.CardNumber == "7755444455551111" {
			return errors.New("daily limit exceeded")
		}
	}

	if !config.DisableFraudDetection {
		if card.CardNumber == "8888888888888888" {
			return errors.New("fraud detected")
		}
	}

	if card.CardNumber == "0000000000000000" {
		return errors.New("card blocked")
	}

	return nil
}

func ValidateAch(ach *model.Ach) error {

	if !achAccountRegex.Match([]byte(ach.AccountNumber)) {
		return errors.New("invalid account number")
	}

	if !achRoutingRegex.Match([]byte(ach.RoutingNumber)) {
		return errors.New("invalid routing number")
	}

	if !achSECRegex.Match([]byte(ach.SECCode)) {
		return errors.New("invalid sec code")
	}

	if !achAccountTypeRegex.Match([]byte(ach.AccountType)) {
		return errors.New("invalid account type")
	}

	return nil
}

func ValidateApplePay(ap *model.ApplePay) error {
	return nil
}

func ValidateGooglePay(gp *model.GooglePay) error {
	return nil
}

func ValidateAddress(address *model.Address) error {

	if !firstNameRegex.Match([]byte(address.FirstName)) {
		return errors.New("invalid first name")
	}

	if !lastNameRegex.Match([]byte(address.LastName)) {
		return errors.New("invalid last name")
	}

	if !postalCodeRegex.Match([]byte(address.PostalCode)) {
		return errors.New("invalid postal code")
	}

	if !cityRegex.Match([]byte(address.City)) {
		return errors.New("invalid city")
	}

	if !addressLineRegex.Match([]byte(address.AddressLine1)) {
		return errors.New("invalid address line")
	}

	if !emailRegex.Match([]byte(address.Email)) {
		return errors.New("invalid email address")
	}

	if !phoneRegex.Match([]byte(address.Phone)) {
		return errors.New("invalid phone number")
	}

	return nil
}

func ValidateXID(xid string) error {

	if !TransactionIdRegex.Match([]byte(xid)) {
		return errors.New("invalid xid")
	}

	return nil
}

func ValidatePaymentMethodType(paymentMethodType string) error {

	if !paymentMethodTypeRegex.Match([]byte(paymentMethodType)) {
		return errors.New("invalid payment method type")
	}

	return nil
}

func ValidatePaymentMethod(paymentMethod *model.PaymentMethod) error {
	if paymentMethod.CreditCard != nil {
		if paymentMethod.Ach != nil || paymentMethod.ApplePay != nil || paymentMethod.GooglePay != nil || paymentMethod.APM != nil {
			return errors.New("only one payment method is allowed")
		}
	}

	if paymentMethod.Ach != nil {
		if paymentMethod.CreditCard != nil || paymentMethod.ApplePay != nil || paymentMethod.GooglePay != nil || paymentMethod.APM != nil {
			return errors.New("only one payment method is allowed")
		}
	}

	if paymentMethod.ApplePay != nil {
		if paymentMethod.Ach != nil || paymentMethod.CreditCard != nil || paymentMethod.GooglePay != nil || paymentMethod.APM != nil {
			return errors.New("only one payment method is allowed")
		}
	}

	if paymentMethod.GooglePay != nil {
		if paymentMethod.Ach != nil || paymentMethod.ApplePay != nil || paymentMethod.CreditCard != nil || paymentMethod.APM != nil {
			return errors.New("only one payment method is allowed")
		}
	}

	if paymentMethod.APM != nil {
		if paymentMethod.Ach != nil || paymentMethod.ApplePay != nil || paymentMethod.GooglePay != nil || paymentMethod.CreditCard != nil {
			return errors.New("only one payment method is allowed")
		}
	}

	return nil
}

func ValidateAmount(amount uint64) error {

	stringFormat := strconv.FormatUint(amount, 10)
	if !amountregex.Match([]byte(stringFormat)) {
		return errors.New("invalid transaction amount")
	}

	return nil
}
