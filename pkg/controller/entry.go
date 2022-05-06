package controller

import (
	"encoding/json"
	"net/http"

	"github.com/ngygbr/mpp/pkg/model"
	"github.com/ngygbr/mpp/pkg/transaction"
	validator "github.com/ngygbr/mpp/pkg/validate"
)

func ProcessCreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var responseModel model.TransactionResponse
	var transactionModel model.Transaction

	err := json.NewDecoder(r.Body).Decode(&transactionModel)
	if err != nil {
		responseModel.SetTransactionResponse(nil, ErrorOccurredCode, err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responseModel)
		return
	}

	//Validate incoming transaction request

	if err = validator.ValidatePaymentMethod(&transactionModel.PaymentMethod); err != nil {
		responseModel.SetTransactionResponse(nil, ErrorOccurredCode, err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responseModel)
		return
	}

	if err = validator.ValidateAddress(&transactionModel.BillingAddress); err != nil {
		responseModel.SetTransactionResponse(nil, ErrorOccurredCode, err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responseModel)
		return
	}

	if err = validator.ValidateAmount(transactionModel.Amount); err != nil {
		responseModel.SetTransactionResponse(nil, ErrorOccurredCode, err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responseModel)
		return
	}

	// Specify the type of the transaction

	if transactionModel.PaymentMethod.CreditCard != nil {
		transactionModel, err = transaction.CCTransaction(transactionModel)
		if err != nil {
			responseModel.SetTransactionResponse(nil, correctResponseCode(err.Error()), err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(responseModel)
			return
		}
	}

	if transactionModel.PaymentMethod.Ach != nil {
		transactionModel, err = transaction.ACHTransaction(transactionModel)
		if err != nil {
			responseModel.SetTransactionResponse(nil, correctResponseCode(err.Error()), err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(responseModel)
			return
		}
	}

	if transactionModel.PaymentMethod.ApplePay != nil {
		transactionModel, err = transaction.ApplePayTransaction(transactionModel)
		if err != nil {
			responseModel.SetTransactionResponse(nil, correctResponseCode(err.Error()), err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(responseModel)
			return
		}
	}

	if transactionModel.PaymentMethod.GooglePay != nil {
		transactionModel, err = transaction.GooglePayTransaction(transactionModel)
		if err != nil {
			responseModel.SetTransactionResponse(nil, correctResponseCode(err.Error()), err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(responseModel)
			return
		}
	}

	responseModel.SetTransactionResponse(&transactionModel, SuccessCode, nil)

	json.NewEncoder(w).Encode(responseModel)
}
