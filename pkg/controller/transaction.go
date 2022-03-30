package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"mpp/pkg/db"
	"mpp/pkg/model"
	validator "mpp/pkg/validate"

	"github.com/gorilla/mux"
)

func ProcessGetTransactionByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	urlParams := mux.Vars(r)
	transactionId := urlParams["id"]

	var response model.ActionResponse

	if err := validator.ValidateXID(transactionId); err != nil {
		response.SetActionResponse(ErrorOccurredCode, err, nil, nil)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	transaction, err := db.GetByID(transactionId)
	if err != nil {
		response.SetActionResponse(ErrorOccurredCode, err, nil, nil)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.SetActionResponse(SuccessCode, nil, &transaction, nil)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		response.SetActionResponse(ErrorOccurredCode, err, nil, nil)
		return
	}
}

func ProcessDeleteTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	urlParams := mux.Vars(r)
	transactionId := urlParams["id"]

	var response model.ActionResponse

	if err := validator.ValidateXID(transactionId); err != nil {
		response.SetActionResponse(ErrorOccurredCode, err, nil, nil)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	err := db.Delete(transactionId)
	if err != nil {
		response.SetActionResponse(ErrorOccurredCode, err, nil, nil)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.SetActionResponse(SuccessCode, nil, nil, nil)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		response.SetActionResponse(ErrorOccurredCode, err, nil, nil)
		return
	}
}

func ProcessGetAllTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var response model.ActionResponse

	transactions, err := db.GetAll()
	if err != nil {
		response.SetActionResponse(ErrorOccurredCode, err, nil, nil)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.SetActionResponse(SuccessCode, nil, nil, &transactions)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		response.SetActionResponse(ErrorOccurredCode, err, nil, nil)
		return
	}
}

func ProcessDeleteAllTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var response model.ActionResponse

	err := db.DeleteAll()
	if err != nil {
		response.SetActionResponse(ErrorOccurredCode, err, nil, nil)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.SetActionResponse(SuccessCode, nil, nil, nil)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		response.SetActionResponse(ErrorOccurredCode, err, nil, nil)
		return
	}
}

func RejectTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	urlParams := mux.Vars(r)
	transactionId := urlParams["id"]

	var response model.ActionResponse

	transaction, err := db.GetByID(transactionId)
	if err != nil {
		response.SetActionResponse(ErrorOccurredCode, err, nil, nil)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	transaction.Status = "rejected"
	transaction.UpdatedAt = time.Now()

	err = db.Update(&transaction)
	if err != nil {
		response.SetActionResponse(ErrorOccurredCode, err, nil, nil)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.SetActionResponse(SuccessCode, nil, nil, nil)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		response.SetActionResponse(ErrorOccurredCode, err, nil, nil)
		return
	}
}

func SettleTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	urlParams := mux.Vars(r)
	transactionId := urlParams["id"]

	var response model.ActionResponse

	transaction, err := db.GetByID(transactionId)
	if err != nil {
		response.SetActionResponse(ErrorOccurredCode, err, nil, nil)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	transaction.Status = "settled"
	transaction.UpdatedAt = time.Now()

	err = db.Update(&transaction)
	if err != nil {
		response.SetActionResponse(ErrorOccurredCode, err, nil, nil)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.SetActionResponse(SuccessCode, nil, nil, nil)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		response.SetActionResponse(ErrorOccurredCode, err, nil, nil)
		return
	}
}

func SetSettledAfter24Hours() {
	for {
		<-time.After(5 * time.Minute)
		//go setSettled()
		go func() {
			transactions, _ := db.GetAll()
			for _, t := range transactions {
				if t.CreatedAt.Add(24 * time.Hour).Before(time.Now()) {
					t.Status = "settled"
					t.UpdatedAt = time.Now()
					err := db.Update(&t)
					if err != nil {
						return
					}
				}
			}
		}()
	}
}
