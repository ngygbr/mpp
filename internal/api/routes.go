package api

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/ngygbr/mpp/pkg/auth"
	configuration "github.com/ngygbr/mpp/pkg/config"
	"github.com/ngygbr/mpp/pkg/controller"
	"github.com/ngygbr/mpp/pkg/healthcheck"
	"net/http"

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
	r.HandleFunc("/login", auth.Login).Methods("GET")
	r.HandleFunc("/encryptcard", controller.ProcessEncryptCard).Methods("POST")

	a.HandleFunc("/transaction/{id}", controller.ProcessGetTransactionByID).Methods("GET")
	a.HandleFunc("/transaction/{id}", controller.ProcessDeleteTransaction).Methods("DELETE")
	a.HandleFunc("/transaction/{id}", preFlight).Methods("OPTIONS")
	a.HandleFunc("/transaction/{id}/settle", controller.SettleTransaction).Methods("GET")
	a.HandleFunc("/transaction/{id}/reject", controller.RejectTransaction).Methods("GET")

	a.HandleFunc("/transactions", controller.ProcessGetAllTransactions).Methods("GET")
	a.HandleFunc("/transactions", preFlight).Methods("OPTIONS")

	a.HandleFunc("/transactions", controller.ProcessDeleteAllTransactions).Methods("DELETE")

	a.HandleFunc("/transaction/creditcard", controller.ProcessCreateTransaction).Methods("POST")

	a.HandleFunc("/transaction/ach", controller.ProcessCreateTransaction).Methods("POST")
	a.HandleFunc("/transaction/applepay", controller.ProcessCreateTransaction).Methods("POST")
	a.HandleFunc("/transaction/googlepay", controller.ProcessCreateTransaction).Methods("POST")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Accept", "X-CSRF-Token", "Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	err := http.ListenAndServe(":"+config.Port, handlers.CORS(originsOk, headersOk, methodsOk)(r))
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

func preFlight(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Request-Headers", "*")
	w.Header().Set("Access-Control-Expose-Headers", "*")

	w.WriteHeader(http.StatusOK)
}
