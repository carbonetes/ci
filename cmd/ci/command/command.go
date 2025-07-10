package command

import (
	"fmt"
	"strings"

	"github.com/carbonetes/ci/internal/constants"
	"github.com/spf13/cobra"
)

func init() {
	// # GENERAL
	root.Flags().BoolP("version", "v", false, "Print the version of Carbonetes CI")

	// # ANALYZER
	root.Flags().StringP("analyzer", "", "", "Analyzer type (jacked or diggity)")
	root.Flags().StringP("scan-type", "", "", fmt.Sprintf("Supported scan types (%s)", strings.Join(constants.SUPPORTED_SCAN_TYPES[:], ", ")))
	root.Flags().StringP("input", "i", "", "Input to be scanned (e.g., image name:tag, filesystem path, tarball path)")
	// # API
	root.Flags().StringP("token", "", "", "Personal Access Token for authentication")
	root.Flags().StringP("plugin-type", "", "", fmt.Sprintf("Supported Plugin types for CI/CD (%s)", strings.Join(constants.SUPPORTED_CICD_PLUGINS[:], ", ")))
	// # FAIL CRITERIA
	root.Flags().StringP("fail-criteria", "", "", fmt.Sprintf("Set the minimum severity level for failing the build based on vulnerability analysis results (build fails if vulnerabilities of this severity or higher are found). Choose Severity:(%s)", strings.Join(constants.FAIL_CRITERIA_SEVERITIES[:], ", ")))
	// # SKIP FAIL
	root.Flags().BoolP("skip-fail", "", false, "Skip failing the build even if any failure criteria or secrets are found.")
	// # FORCE DB UPDATE
	root.Flags().BoolP("force-db-update", "", false, "Force update of the database even if it is already up to date. This will re-download the database files and update the local database.")
	// # ENVIRONMENT TYPE
	root.Flags().StringP("environment-type", "", "", fmt.Sprintf("Supported environment types (%s)", strings.Join(constants.SUPPORTED_ENVIRONMENT_TYPE[:], ", ")))

	// # HELP
	root.PersistentFlags().BoolP("help", "h", false, "")
	root.PersistentFlags().Lookup("help").Hidden = true
	root.SetHelpCommand(&cobra.Command{Hidden: true})
}

func Run() error {
	if err := root.Execute(); err != nil {
		return err
	}
	return nil
}
