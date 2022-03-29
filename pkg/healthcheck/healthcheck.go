package healthcheck

import (
	"fmt"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	_, err := fmt.Fprintf(w, "API is up and running...")
	if err != nil {
		panic("healthcheck failed")
	}
}
