package api

const (
	// Encrypted URLs (base64-encoded AES-GCM ciphertext)
	encryptedLocalhost = "ENCRYPTED_BASE64_STRING_FOR_LOCALHOST"
	encryptedTapp      = "ENCRYPTED_BASE64_STRING_FOR_TAPP"
	encryptedProd      = "ENCRYPTED_BASE64_STRING_FOR_PROD"
)

// EnvironmentTypeSelector returns the decrypted URL based on the type
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
		encryptedURL = encryptedTapp
	}

	decryptedURL, err := util.decryptAESGCM(encryptedURL)
	if err != nil {
		return "", err
	}

	return decryptedURL, nil
}
