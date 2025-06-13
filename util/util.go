package util

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
)

var (
	raw     = []byte("!CarbonetesContinuousIntegration")
	refined = []byte("!Carbo-CI-iv")
)

// EncryptAESGCM encrypts plaintext using fixed key and IV
func EncryptAESGCM(plainText string) (cipherBase64 string, err error) {
	block, err := aes.NewCipher(raw)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	if len(refined) != aesGCM.NonceSize() {
		return "", fmt.Errorf("invalid IV size")
	}

	cipherText := aesGCM.Seal(nil, refined, []byte(plainText), nil)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// DecryptAESGCM decrypts base64 ciphertext using fixed key and IV
func DecryptAESGCM(cipherBase64 string) (string, error) {
	cipherText, err := base64.StdEncoding.DecodeString(cipherBase64)
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

	if len(refined) != aesGCM.NonceSize() {
		return "", errors.New("invalid IV size")
	}

	plainText, err := aesGCM.Open(nil, refined, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}
