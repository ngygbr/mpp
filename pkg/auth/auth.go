package auth

import (
	"encoding/json"
	"net/http"
	"time"

	configuration "mpp/pkg/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

type User struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

var config = configuration.GetConfig()

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	token, err := generateJWT(user.Name)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	user.Token = token

	json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func generateJWT(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	if username == "" {
		return "", errors.New("invalid username")
	}

	claims["user"] = username
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

	tokenString, err := token.SignedString([]byte(config.SignKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
