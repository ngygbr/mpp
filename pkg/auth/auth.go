package auth

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"time"

	configuration "mpp/pkg/config"

	"github.com/dgrijalva/jwt-go"
)

type user struct {
	Token string `json:"token"`
}

var config = configuration.GetConfig()

func Login(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user user
	token, err := generateJWT()
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

func generateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	hexadec := hex.EncodeToString(bytes)

	claims["user"] = hexadec
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

	tokenString, err := token.SignedString([]byte(config.SignKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
