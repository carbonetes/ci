package util

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

var raw = []byte("the-32-byte-long-wordss-xxxx!!")

func decryptAESGCM(cipherTextBase64 string) (string, error) {
	cipherText, err := base64.StdEncoding.DecodeString(cipherTextBase64)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(raw)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	if len(cipherText) < aesGCM.NonceSize() {
		return "", errors.New("ciphertext too short")
	}

	nonce := cipherText[:aesGCM.NonceSize()]
	ciphertext := cipherText[aesGCM.NonceSize():]

	plainText, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}
