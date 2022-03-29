package controller

import (
	"encoding/json"
	"mock-paymentprocessor/pkg/model"
	transaction2 "mock-paymentprocessor/pkg/transaction"
	"mock-paymentprocessor/pkg/transaction/types"
	validate "mock-paymentprocessor/pkg/validate"
	"net/http"
)

func ProcessCreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var response model.TransactionResponse
	var transaction model.Transaction

	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		response.SetTransactionResponse(nil, transaction2.ErrorOccurredCode, err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	if err = validate.ValidatePaymentMethod(&transaction.PaymentMethod); err != nil {
		response.SetTransactionResponse(nil, transaction2.ErrorOccurredCode, err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	if transaction.PaymentMethod.CreditCard != nil {
		transaction, err = types.CCTransaction(transaction)
		if err != nil {
			response.SetTransactionResponse(nil, correctResponseCode(err.Error()), err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	if transaction.PaymentMethod.Ach != nil {
		transaction, err = types.ACHTransaction(transaction)
		if err != nil {
			response.SetTransactionResponse(nil, correctResponseCode(err.Error()), err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	if transaction.PaymentMethod.ApplePay != nil {
		transaction, err = types.ApplePayTransaction(transaction)
		if err != nil {
			response.SetTransactionResponse(nil, correctResponseCode(err.Error()), err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	if transaction.PaymentMethod.GooglePay != nil {
		transaction, err = types.GooglePayTransaction(transaction)
		if err != nil {
			response.SetTransactionResponse(nil, correctResponseCode(err.Error()), err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	if transaction.PaymentMethod.APM != nil {
		transaction, err = types.APMTransaction(transaction)
		if err != nil {
			response.SetTransactionResponse(nil, correctResponseCode(err.Error()), err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	if err != nil {
		response.SetTransactionResponse(nil, transaction2.ErrorOccurredCode, err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.SetTransactionResponse(&transaction, transaction2.SuccessCode, nil)

	json.NewEncoder(w).Encode(response)
	if err != nil {
		response.SetTransactionResponse(nil, transaction2.ErrorOccurredCode, err)
		return
	}
}

func correctResponseCode(message string) uint64 {
	switch message {
		case "fraud detected":
			return transaction2.FraudDetectedCode
		case "limit exceeded":
			return transaction2.LimitExceededCode
		case "daily limit exceeded":
			return transaction2.DailyLimitExceededCode
		case "card blocked":
			return transaction2.CardBlockedCode
		case "success":
			return transaction2.SuccessCode
	}
	return transaction2.ErrorOccurredCode
}
