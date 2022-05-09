package controller

import (
	"encoding/json"
	crypter "github.com/ngygbr/mpp/pkg/crypto"
	"github.com/ngygbr/mpp/pkg/model"
	"net/http"
)

type EncryptedData struct {
	Card          *model.CreditCard `json:"credit_card,omitempty"`
	EncryptionKey string            `json:"encryption_key"`
	EncryptedCard []byte            `json:"encrypted_card"`
}

func ProcessEncryptCard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var encryptedData EncryptedData

	err := json.NewDecoder(r.Body).Decode(&encryptedData)
	if err != nil {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	ciphertext, encKey, err := crypter.EncryptCard(encryptedData.Card, encryptedData.EncryptionKey)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	encryptedData.EncryptedCard = []byte(ciphertext)
	encryptedData.EncryptionKey = encKey
	encryptedData.Card = nil

	err = json.NewEncoder(w).Encode(encryptedData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
