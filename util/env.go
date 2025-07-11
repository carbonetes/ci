package util

import (
	"os"

	"github.com/carbonetes/ci/internal/constants"
	"github.com/carbonetes/ci/internal/log"
)

const (
	LOCALHOSTURL = "LvP/3uiWyFaw7vyI97Xkg2Od4FgGGguDZ+K9h9r1I7yAH/n7tRk="
	TAPPURL      = "LvP/3qGDyBWr6POQsrvnnnfNsRpVWkGq2PUh8klXC2+ksb06MLmhyISj3ii+bE4V"
	PRODURL      = "" // Not Available
)

func EnvironmentTypeSelector(environmentType int) (string, error) {
	var encryptedURL string

	switch environmentType {
	case 0:
		encryptedURL = LOCALHOSTURL
	case 1:
		encryptedURL = TAPPURL
	case 2:
		encryptedURL = PRODURL
	default:
		encryptedURL = TAPPURL
	}

	return DecryptAESGCM(encryptedURL)
}

func GetEnvironmentType(environmentType string) (envType int) {
	switch environmentType {
	case "localhost":
		envType = 0
	case "test":
		envType = 1
	case "production":
		envType = 2
	default:
		log.Printf("%v: Invalid environment type %s. Supported environment types are: %v", constants.CI_FAILURE, environmentType, constants.SUPPORTED_ENVIRONMENT_TYPE)
		os.Exit(1)
	}
	return envType

}
