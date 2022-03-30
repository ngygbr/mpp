package api

import (
	"fmt"
	"net/http"

	"mpp/pkg/auth"
	configuration "mpp/pkg/config"
	"mpp/pkg/controller"
	"mpp/pkg/healthcheck"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

var config = configuration.GetConfig()

func Init() error {

	r := mux.NewRouter()
	a := r.PathPrefix("/api").Subrouter()
	a.Use(authMiddleware)

	r.HandleFunc("/healthcheck", healthcheck.HealthCheck).Methods("GET")
	r.HandleFunc("/login", auth.Login).Methods("POST")

	a.HandleFunc("/transaction/{id}", controller.ProcessGetTransactionByID).Methods("GET")
	a.HandleFunc("/transaction/{id}", controller.ProcessDeleteTransaction).Methods("DELETE")
	a.HandleFunc("/transaction/{id}/settle", controller.SettleTransaction).Methods("POST")
	a.HandleFunc("/transaction/{id}/reject", controller.RejectTransaction).Methods("POST")
	a.HandleFunc("/transactions", controller.ProcessGetAllTransactions).Methods("GET")
	a.HandleFunc("/transactions", controller.ProcessDeleteAllTransactions).Methods("DELETE")

	a.HandleFunc("/transaction/creditcard", controller.ProcessCreateTransaction).Methods("POST")
	a.HandleFunc("/transaction/ach", controller.ProcessCreateTransaction).Methods("POST")
	a.HandleFunc("/transaction/applepay", controller.ProcessCreateTransaction).Methods("POST")
	a.HandleFunc("/transaction/googlepay", controller.ProcessCreateTransaction).Methods("POST")

	err := http.ListenAndServe(":"+config.Port, r)
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
