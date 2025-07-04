package command

import (
	"fmt"
	"strings"

	"github.com/carbonetes/ci/internal/constants"
)

func init() {
	// # GENERAL
	root.Flags().BoolP("version", "v", false, "Print the version of Carbonetes CI")

	// # ANALYZER
	root.Flags().StringP("analyzer", "", "", "Analyzer type (jacked or diggity)")
	// # API
	root.Flags().StringP("token", "", "", "Personal Access Token for authentication")
	root.Flags().StringP("plugin-type", "", "", fmt.Sprintf("Supported Plugin types for CI/CD (%s)", strings.Join(constants.SUPPORTED_CICD_PLUGINS[:], ", ")))
	// # FAIL CRITERIA
	root.Flags().StringP("fail-criteria", "", "", fmt.Sprintf("Set the minimum severity level for failing the build based on vulnerability analysis results (build fails if vulnerabilities of this severity or higher are found). Choose Severity:(%s)", strings.Join(constants.FAIL_CRITERIA_SEVERITIES[:], ", ")))
	// # SKIP FAIL
	root.Flags().BoolP("skip-fail", "", false, "Skip failing the build even if vulnerabilities, secrets, or any input error are found during the analysis")
}

func Run() error {
	if err := root.Execute(); err != nil {
		return err
	}
	return nil
}
