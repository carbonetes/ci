package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/carbonetes/ci/internal/constants"
	"github.com/carbonetes/ci/internal/log"
)

// apiRequest performs an HTTP POST request with JSON payload.
func apiRequest(payload any, url string) (*http.Response, []byte) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("%v Unable to process the payload from the selected environment. Please report this issue.", constants.CI_FAILURE)
		os.Exit(1)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("%v Unable to access the selected environment. Please report this issue.", constants.CI_FAILURE)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("%v Unable to read the return response body from the selected environment. Please report this issue.", constants.CI_FAILURE)
		os.Exit(1)
	}

	return resp, body
}
