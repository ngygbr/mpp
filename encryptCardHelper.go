package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mpp/pkg/model"
)

/*
EXAMPLE ENCRYPTED CARDS TO USE IN APPLE PAY / GOOGLE PAY REQUEST

key: c5a981e16742f8bfb7056bcdc2817bac512e3849394f65895cbbe88471951aa8
encrypted : 4d91fad8c3e7ae5856e23559a6b772fde039e910efb654aaf1328e54b1beefff59c33dea0b080cbea4aa4f407a607cedf8a78e5027584d1030547336f419456a282c4ff63cd331a2d5a512f4107a46c733032ec3f34331f83b973760a4114c0a277042312491c8e93d805a24653838e35de2e824fa2d
----------------------------------------------------

key: 484d1cf96c8409e02c4c71276f265b65b8329bc1f8438cf66c08c975a7d4b84a
encrypted : 38041f2368c5118806ed23951fe0f166e2f64099b6f6be495f5fbb248a154a0bf11e11a4bc47749d3e589eaeb59b428ae6b04ea1563140d5ef2118f623da8fdd06ed4c323560303d7ff1d15a5aacf6e93d9083fa21903ab5de65adbc3667a08cbe2cecb5beebbbe11cbdbebccad0d7e91d8f561f02466ffb70
----------------------------------------------------

key: 88c70221c13df0abd0e1ac9a1cb18926c062ca0906dc1b0a0ab4c2d3b17a0756
encrypted : 1b67bfe71dd9caf40d86a2e6bdce33a2dc6fdfac82dee3104424e8d9b58b606af667c9080f4a0eeff9b1a0eae3d2abf2d755a9dcff4a0c057a86d24c2da567c0739552c521b5b2a7911f7909c8fd953b94ea018d2236ef9bc4831c2a8212a6b10a7b4ada19d0a2335a63705aded0d5fd2629f5cd61ac
----------------------------------------------------

 */

/*
func main() {

	//Example Cards to encrypt
	cc1 := model.CreditCard{
		CardNumber:     "5444166444123444",
		HolderName:     "Test Man",
		ExpirationDate: "01/30",
		CVC:            "123",
	}
	cc2 := model.CreditCard{
		CardNumber:     "4333111222111333",
		HolderName:     "Decrypt Guy",
		ExpirationDate: "09/28",
		CVC:            "444",
	}
	cc3 := model.CreditCard{
		CardNumber:     "5444166444123444",
		HolderName:     "Test Man",
		ExpirationDate: "01/30",
		CVC:            "123",
	}

	//random 32 byte key for AES-256
	bytes1 := make([]byte, 32)
	if _, err := rand.Read(bytes1); err != nil {
		panic(err.Error())
	}
	bytes2 := make([]byte, 32)
	if _, err := rand.Read(bytes2); err != nil {
		panic(err.Error())
	}
	bytes3 := make([]byte, 32)
	if _, err := rand.Read(bytes3); err != nil {
		panic(err.Error())
	}

	//generated keys for encrypting
	key1 := hex.EncodeToString(bytes1)
	key2 := hex.EncodeToString(bytes2)
	key3 := hex.EncodeToString(bytes3)

	//Encrypted data with keys
	encrypted1 := encrypt(cc1, key1)
	fmt.Printf("key: %s\n", key1)
	fmt.Printf("encrypted : %s\n", encrypted1)
	fmt.Println("----------------------------------------------------")
	encrypted2 := encrypt(cc2, key2)
	fmt.Printf("key: %s\n", key2)
	fmt.Printf("encrypted : %s\n", encrypted2)
	fmt.Println("----------------------------------------------------")
	encrypted3 := encrypt(cc3, key3)
	fmt.Printf("key: %s\n", key3)
	fmt.Printf("encrypted : %s\n", encrypted3)
	fmt.Println("----------------------------------------------------")
}
 */

func encrypt(cardToEncrypt model.CreditCard, keyString string) (encryptedString string) {

	marshalledCard, err := json.Marshal(cardToEncrypt)
	if err != nil {
		log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	}

	key, _ := hex.DecodeString(keyString)
	plaintext := marshalledCard

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext)
}
