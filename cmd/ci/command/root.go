package command

import (
	"os"

	"github.com/carbonetes/ci/cmd/ci/build"
	"github.com/carbonetes/ci/internal/constants"
	"github.com/carbonetes/ci/internal/helper"
	"github.com/carbonetes/ci/internal/log"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "ci",
	Args:  cobra.MaximumNArgs(1),
	Short: "Carbonetes CI",
	Long:  `Carbonetes Continuous Integration`,
	Run:   rootCmd,
}

func rootCmd(c *cobra.Command, args []string) {
	versionArg, _ := c.Flags().GetBool("version")
	if versionArg {
		log.Print(build.GetBuild().Version)
	}

	// Retrieve flags
	analyzer, _ := c.Flags().GetString("analyzer")
	token, _ := c.Flags().GetString("token")
	pluginType, _ := c.Flags().GetString("plugin-type")
	failCriteria, _ := c.Flags().GetString("fail-criteria")
	skipFail, _ := c.Flags().GetBool("skip-fail")

	// # INPUT CHECKING

	// ## SKIP FAIL FLAG
	if skipFail {
		log.Warnf("%v Skip fail is ENABLED!.", constants.CI_WARNING)
	}

	// ## JACKED & DIGGITY FLAGS
	if len(analyzer) > 0 {
		switch analyzer {
		case constants.JACKED:
			if len(failCriteria) == 0 {
				log.Fatalf("%v: Fail criteria is supported for jacked analyzer", constants.CI_FAILURE)
				os.Exit(1)
			}
		case constants.DIGGITY:
			if len(failCriteria) > 0 {
				log.Fatalf("%v: Fail criteria is not supported for diggity analyzer", constants.CI_FAILURE)
				os.Exit(1)
			}
		default:
			log.Fatalf("%v: Invalid analyzer type %s", constants.CI_FAILURE, analyzer)
			os.Exit(1)
		}
	}

	// ## API TAGS: TOKEN & PLUGIN TYPE FLAGS
	if len(token) == 0 {
		log.Fatalf("%v: No token provided. Use --token flag to provide a token.", constants.CI_FAILURE)
		os.Exit(1)
	}
	if len(pluginType) == 0 {
		log.Fatalf("%v: No plugin type provided. Use --plugin-type flag to provide a plugin type.", constants.CI_FAILURE)
		os.Exit(1)
	}
	if len(pluginType) > 0 && !helper.Contains(constants.SUPPORTED_CICD_PLUGINS[:], pluginType) {
		log.Fatalf("%v: Invalid plugin type %s. Supported types are: %v", constants.CI_FAILURE, pluginType, constants.SUPPORTED_CICD_PLUGINS)
		os.Exit(1)
	}

	// ## FAIL CRITERIA FLAG
	if len(failCriteria) > 0 && !helper.Contains(constants.FAIL_CRITERIA_SEVERITIES[:], failCriteria) {
		log.Fatalf("%v: Invalid fail criteria %s. Supported criteria are: %v", constants.CI_FAILURE, failCriteria, constants.FAIL_CRITERIA_SEVERITIES)
		os.Exit(1)
	}

}
