package api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"mock-paymentprocessor/pkg/auth"
	"mock-paymentprocessor/pkg/config"
	"mock-paymentprocessor/pkg/controller"
	"net/http"

	"mock-paymentprocessor/pkg/healthcheck"
	"mock-paymentprocessor/pkg/transaction"

	"github.com/gorilla/mux"
)

var config = utils.GetConfig()

func Init() error {

	r := mux.NewRouter()
	a := r.PathPrefix("/api").Subrouter()
	a.Use(authMiddleware)

	r.HandleFunc("/healthcheck", healthcheck.HealthCheck).Methods("GET")
	r.HandleFunc("/login", auth.Login).Methods("POST")

	a.HandleFunc("/transaction/{id}", transaction.ProcessGetTransactionByID).Methods("GET")
	a.HandleFunc("/transaction/{id}", transaction.ProcessDeleteTransaction).Methods("DELETE")
	a.HandleFunc("/transaction/{id}/settle", transaction.SettleTransaction).Methods("POST")
	a.HandleFunc("/transaction/{id}/reject", transaction.RejectTransaction).Methods("POST")
	a.HandleFunc("/transactions", transaction.ProcessGetAllTransactions).Methods("GET")
	a.HandleFunc("/transactions", transaction.ProcessDeleteAllTransactions).Methods("DELETE")

	a.HandleFunc("/transaction/creditcard", controller.ProcessCreateTransaction).Methods("POST")
	a.HandleFunc("/transaction/ach", controller.ProcessCreateTransaction).Methods("POST")
	a.HandleFunc("/transaction/applepay", controller.ProcessCreateTransaction).Methods("POST")
	a.HandleFunc("/transaction/googlepay", controller.ProcessCreateTransaction).Methods("POST")

	err := http.ListenAndServe(":" + config.Port, r)
	if err != nil {
		return err
	}

	return nil
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Authorization"] != nil {
			token, err := jwt.Parse(r.Header["Authorization"][0], func(token *jwt.Token) (interface{}, error) {

				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("error signing method")
				}

				return []byte(config.SignKey), nil

			})
			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				next.ServeHTTP(w, r)
			}

		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}
