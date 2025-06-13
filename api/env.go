package api

import (
	"errors"

	"github.com/carbonetes/ci/util"
)

const (
	// Encrypted base64-encoded (nonce + ciphertext)
	encryptedLocalhost = "ENCRYPTED_BASE64_FOR_LOCALHOST"
	encryptedTapp      = "ENCRYPTED_BASE64_FOR_TAPP"
	encryptedProd      = "ENCRYPTED_BASE64_FOR_PROD"
)

func EnvironmentTypeSelector(environmentType int) (string, error) {
	var encryptedURL string

	switch environmentType {
	case 0:
		encryptedURL = encryptedLocalhost
	case 1:
		encryptedURL = encryptedTapp
	case 2:
		encryptedURL = encryptedProd
	default:
		return "", errors.New("invalid environment type")
	}

	return util.DecryptAESGCM(encryptedURL)
}
