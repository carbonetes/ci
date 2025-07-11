package oss

import (
	"os"

	"github.com/carbonetes/ci/api"
	"github.com/carbonetes/ci/cmd/ci/oss/diggity"
	"github.com/carbonetes/ci/internal/constants"
	"github.com/carbonetes/ci/internal/helper"
	"github.com/carbonetes/ci/internal/log"
	"github.com/carbonetes/ci/internal/presenter"
	"github.com/carbonetes/ci/pkg/types"
)

func Run(parameters types.Parameters) {

	// Start Personal Access Token Validation
	api.PersonalAccessToken(parameters.Token, parameters.PluginType, parameters.EnvType)

	switch parameters.ScanType {
	// FILE SYSTEM / DIRECTORY
	case constants.FILE_SYSTEM:
		if found, _ := helper.IsDirExists(parameters.Input); !found {
			log.Printf("%v: Input path '%s' does not exist or is not a directory.", constants.CI_FAILURE, parameters.Input)
			os.Exit(1)
		}
		parameters.Diggity.ScanType = 3
	// TARBALL / TAR FILE
	case constants.TARBALL:
		if found, _ := helper.IsFileExists(parameters.Input); !found {
			log.Printf("%v: Input tar file '%s' does not exists.", constants.CI_FAILURE, parameters.Input)
			os.Exit(1)
		}
		parameters.Diggity.ScanType = 2
	case constants.IMAGE:
		if len(parameters.Input) > 0 {
			parameters.Diggity.ScanType = 1
			parameters.Input = helper.FormatImage(parameters.Input)
		} else {
			log.Printf("%v: Input path is required for image scan type.", constants.CI_FAILURE)
			os.Exit(1)
		}
	default:
		log.Printf("%v: Unsupported scan type '%s'. Supported scan types are: %v", constants.CI_FAILURE, parameters.ScanType, constants.SUPPORTED_SCAN_TYPES)
		os.Exit(1)
	}

	presenter.DisplayInput(parameters)

	bom, secrets := diggity.Analyze(parameters)

	// # Save Analysis Result
	api.SavePluginRepository(bom, parameters.Input, parameters.PluginType, parameters.Duration, parameters.EnvType, parameters.AnalysisType, secrets)
}
