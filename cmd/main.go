package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func encryptAESGCM(plainText string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherText := aesGCM.Seal(nonce, nonce, []byte(plainText), nil)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func main() {
	key := []byte("the-32-byte-long-wordss-xxxx!!") // Same as in `crypto.go`

	urls := []string{
		"http://localhost:3001/",
		"https://tent-api.carbonetes.com/",
		"https://prod.carbonetes.com/",
	}

	for _, url := range urls {
		enc, err := encryptAESGCM(url, key)
		if err != nil {
			panic(err)
		}
		fmt.Println(enc)
	}
}
