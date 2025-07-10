package api

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/CycloneDX/cyclonedx-go"
	"github.com/carbonetes/ci/internal/constants"
	"github.com/carbonetes/ci/internal/log"
	"github.com/carbonetes/ci/util"
	"github.com/carbonetes/diggity/pkg/types"
)

// SavePluginRepository submits SBOM and metadata to Carbonetes API.
func SavePluginRepository(bom *cyclonedx.BOM, repoName, pluginName string, start time.Time, environmentType int, analysisType int, secrets []types.Secret) {

	url, err := util.EnvironmentTypeSelector(environmentType)
	if err != nil {
		log.Fatalf("%v: Sync: Something went wrong on getting environment type. Please report this issue.", constants.CI_FAILURE)
		os.Exit(1)
	}

	analysis := util.AnalysisTypeSelector(analysisType)

	var bomBytes []byte
	if bom != nil {
		bomBytes, err = json.Marshal(bom)
		if err != nil {
			log.Fatalf("%v: Sync: Something went wrong on processing packages. Please report this issue.", constants.CI_FAILURE)
			os.Exit(1)
		}
	}
	var secretBytes []byte
	if len(secrets) > 0 {
		secretBytes, err = json.Marshal(secrets)
		if err != nil {
			log.Fatalf("%v: Sync: Something went wrong on processing secrets. Please report this issue.", constants.CI_FAILURE)
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
		log.Fatalf("%v: Sync: Fail to process response body from the selected environment. Please report this issue.", constants.CI_FAILURE)
		os.Exit(1)
	}

	if resp.StatusCode != 200 {
		log.Fatalf("%v: Syncing Analysis Result Failed. Please report this issue.", constants.CI_FAILURE)
		os.Exit(1)
	}
}
