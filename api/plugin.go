package api

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/CycloneDX/cyclonedx-go"
	"github.com/carbonetes/ci/util"
	"github.com/carbonetes/diggity/pkg/types"
)

// SavePluginRepository submits SBOM and metadata to Carbonetes API.
func SavePluginRepository(bom *cyclonedx.BOM, repoName, pluginName string, start time.Time, environmentType int, analysisType int, secrets []types.Secret) {

	url, err := util.EnvironmentTypeSelector(environmentType)
	if err != nil {
		fmt.Println("Failed to parse response:", err)
		os.Exit(1)
	}

	analysis := util.AnalysisTypeSelector(analysisType)

	var bomBytes []byte
	if bom != nil {
		bomBytes, err = json.Marshal(bom)
		if err != nil {
			fmt.Println("Failed to marshal BOM:", err)
			os.Exit(1)
		}
	}
	var secretBytes []byte
	if len(secrets) > 0 {
		secretBytes, err = json.Marshal(secrets)
		if err != nil {
			fmt.Println("Failed to marshal BOM:", err)
			os.Exit(1)
		}
	}

	payload := map[string]interface{}{
		"repoName":              repoName,
		"personalAccessTokenId": tokenId,
		"pluginName":            pluginName,
		"bom":                   bomBytes,
		"secrets":               secretBytes,
		"duration":              fmt.Sprintf("%.2f", time.Since(start).Seconds()),
	}

	resp, body := apiRequest(payload, fmt.Sprintf("%sintegrations/"+analysis+"/plugin/save", url))

	var result PluginRepo
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Failed to parse response:", err)
		os.Exit(1)
	}

	if resp.StatusCode != 200 {
		fmt.Println("Status Code:", resp.StatusCode)
		fmt.Println("Response Body:", string(body))
		os.Exit(1)
	}
}
