package ci

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/CycloneDX/cyclonedx-go"
)

// SavePluginRepository submits SBOM and metadata to Carbonetes API.
func SavePluginRepository(bom *cyclonedx.BOM, repoName, pluginName string, start time.Time) {
	var bomJSONString string
	if bom != nil {
		bomBytes, err := json.Marshal(bom)
		if err != nil {
			fmt.Println("Failed to marshal BOM:", err)
			os.Exit(1)
		}
		bomJSONString = string(bomBytes)
	}

	payload := map[string]interface{}{
		"repoName":              repoName,
		"personalAccessTokenId": tokenId,
		"pluginName":            pluginName,
		"bom":                   bomJSONString,
		"duration":              fmt.Sprintf("%.2f", time.Since(start).Seconds()),
	}

	resp, body := apiRequest(payload, saveURL)

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
