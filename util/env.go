package util

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
