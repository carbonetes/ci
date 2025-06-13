package ci

import (
	"encoding/json"
	"fmt"
	"os"

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

	payload := map[string]string{
		"token":      token,
		"pluginType": pluginType,
	}

	resp, body := apiRequest(payload, fmt.Sprintf("%s/personal-access-token/is-expired", url))

	var result TokenCheckResponse
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Failed to parse response:", err)
		os.Exit(1)
	}

	for _, p := range result.Permissions {
		if p.Label == "Pipelines" {
			for _, lp := range p.Permissions {
				if lp == "write" {
					permitted = true
					break
				}
			}
		}
	}

	if !permitted {
		fmt.Println("Status Code:", 401)
		fmt.Println("Error: You do not have pipeline write permission.")
		os.Exit(1)
	}

	if resp.StatusCode != 200 {
		fmt.Println("Status Code:", resp.StatusCode)
		fmt.Println("Response Body:", string(body))
		os.Exit(1)
	}

	tokenId = result.PersonalAccessTokenId
	if tokenId == "" {
		fmt.Println("Status Code:", resp.StatusCode)
		fmt.Println("Error: Unable to fetch token id.")
		os.Exit(1)
	}
}
