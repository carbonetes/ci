package api

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/carbonetes/ci/internal/log"

	"github.com/carbonetes/ci/util"
)

var tokenId = "0"
var permitted = false

// PersonalAccessToken checks token permissions and sets global token ID.
func PersonalAccessToken(token, pluginType string, environmentType int) {

	url, err := util.EnvironmentTypeSelector(environmentType)
	if err != nil {
		fmt.Println("Failed to parse response:", err)
		os.Exit(1)
	}

	// Payload
	payload := map[string]string{
		"token":      token,
		"pluginType": pluginType,
	}

	// Perform HTTP POST request
	resp, body := apiRequest(payload, fmt.Sprintf("%spersonal-access-token/is-expired", url))
	// ---------------

	if resp.StatusCode != 200 {
		var appError ApplicationErrorResponse
		if err := json.Unmarshal(body, &appError); err != nil {
			log.Fatal("Failed to parse response:", err)
			os.Exit(1)
		}
		log.Print("Error: ", appError.Message)
		os.Exit(1)
	}
	// Unmarshal the body into the struct
	var result TokenCheckResponse
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatal("Failed to parse response:", err)
		os.Exit(1)
	}

	for _, p := range result.Permissions {
		if p.Label == "Pipelines" {
			for _, lp := range p.Permissions {
				if lp == "write" {
					permitted = true
				}
			}
		}
	}

	if !permitted {
		log.Fatal("Error: You do not have pipeline write permission.")
		os.Exit(1)
	}

	tokenId = result.PersonalAccessTokenId
	if result.PersonalAccessTokenId == "" {
		log.Fatal("Status Code:", resp.StatusCode)
		log.Fatal("Error: Unable to fetch token id.")
		os.Exit(1)
	}

}
