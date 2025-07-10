package api

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/carbonetes/ci/internal/log"

	"github.com/carbonetes/ci/internal/constants"
	"github.com/carbonetes/ci/util"
)

var tokenId = "0"
var permitted = false

// PersonalAccessToken checks token permissions and sets global token ID.
func PersonalAccessToken(token, pluginType string, environmentType int) {

	url, err := util.EnvironmentTypeSelector(environmentType)
	if err != nil {
		log.Fatalf("%v: EnvironmentTypeSelector Error. Please report this issue.", constants.CI_FAILURE)
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
			log.Fatalf("%v: Fail to process response body from the selected environment. Please report this issue.", constants.CI_FAILURE)
			os.Exit(1)
		}
		log.Fatalf("%v: %v", constants.CI_FAILURE, appError.Message)
		os.Exit(1)
	}
	// Unmarshal the body into the struct
	var result TokenCheckResponse
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("%v: Fail to process response body from the selected environment. Please report this issue.", constants.CI_FAILURE)
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
		log.Fatalf("%v: Token does not have PIPELINE Write Permission.", constants.CI_FAILURE)
		os.Exit(1)
	}

	tokenId = result.PersonalAccessTokenId
	if result.PersonalAccessTokenId == "" {
		log.Fatalf("%v: Something went wrong from getting token details. Please report this issue.", constants.CI_FAILURE)
		os.Exit(1)
	}

}
