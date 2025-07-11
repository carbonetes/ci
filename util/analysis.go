package util

import (
	"os"

	"github.com/carbonetes/ci/internal/constants"
	"github.com/carbonetes/ci/internal/log"
)

const (
	BOM  = "bom"
	VULN = "vuln"
)

func AnalysisTypeSelector(analysisType int) string {
	var analysisTypeStr string

	switch analysisType {
	case 0:
		analysisTypeStr = BOM
	case 1:
		analysisTypeStr = VULN
	default:
		analysisTypeStr = BOM
	}

	return analysisTypeStr
}

func GetAnalysisType(analyzer string) (analysisType int) {

	switch analyzer {
	case "diggity":
		analysisType = 0
	case "jacked":
		analysisType = 1
	default:
		log.Printf("%v: Invalid environment type %s. Supported environment types are: %v", constants.CI_FAILURE, analysisType, constants.SUPPORTED_ANALYZERS)
		os.Exit(1)
	}
	return analysisType
}
