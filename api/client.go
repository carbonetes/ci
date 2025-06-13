package ci

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// apiRequest performs an HTTP POST request with JSON payload.
func apiRequest(payload any, url string) (*http.Response, []byte) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return resp, body
}
