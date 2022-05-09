package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"github.com/ngygbr/mpp/pkg/model"
	validator "github.com/ngygbr/mpp/pkg/validate"
	"io"
	"log"
)

func generateKey() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	key := hex.EncodeToString(bytes)

	return key, nil
}

func EncryptCard(cardToEncrypt *model.CreditCard, keyString string) (string, string, error) {

	if err := validator.ValidateCreditCard(cardToEncrypt); err != nil {
		return "", "", err
	}

	marshalledCard, err := json.Marshal(cardToEncrypt)
	if err != nil {
		log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	}

	key, _ := hex.DecodeString(keyString)
	plaintext := marshalledCard

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", "", err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return string(ciphertext), keyString, nil
}

func DecryptCard(encryptedString string, keyString string) (string, error) {

	key, err := hex.DecodeString(keyString)
	if err != nil {
		return "", err
	}

	enc, err := hex.DecodeString(encryptedString)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()

	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

