package command

import (
	"os"

	"github.com/carbonetes/ci/cmd/ci/build"
	"github.com/carbonetes/ci/cmd/ci/oss"
	"github.com/carbonetes/ci/internal/constants"
	"github.com/carbonetes/ci/internal/helper"
	"github.com/carbonetes/ci/internal/log"
	"github.com/carbonetes/ci/internal/presenter"
	"github.com/carbonetes/ci/pkg/types"
	"github.com/carbonetes/ci/util"
	"github.com/spf13/cobra"

	diggity "github.com/carbonetes/diggity/pkg/types"
	jacked "github.com/carbonetes/jacked/pkg/types"
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
		return
	}

	// Retrieve flags
	analyzer, _ := c.Flags().GetString("analyzer")
	scanType, _ := c.Flags().GetString("scan-type")
	input, _ := c.Flags().GetString("input")
	token, _ := c.Flags().GetString("token")
	pluginType, _ := c.Flags().GetString("plugin-type")
	failCriteria, _ := c.Flags().GetString("fail-criteria")
	skipFail, _ := c.Flags().GetBool("skip-fail")
	forceDbUpdate, _ := c.Flags().GetBool("force-db-update")
	environmentType, _ := c.Flags().GetString("environment-type")

	// # INPUT CHECKING
	if len(input) == 0 {
		c.Help()
		return
	}

	// ## SKIP FAIL FLAG
	if skipFail {
		presenter.DisplaySkipFail()
	}

	// ## JACKED & DIGGITY FLAGS
	if len(analyzer) > 0 {
		switch analyzer {
		case constants.JACKED:
			// # FAIL CRITERIA FLAG
			if len(failCriteria) == 0 {
				log.Fatalf("%v: Fail criteria is supported for jacked analyzer", constants.CI_FAILURE)
				os.Exit(1)
			}

			if len(failCriteria) > 0 && !helper.Contains(constants.FAIL_CRITERIA_SEVERITIES[:], failCriteria) {
				log.Fatalf("%v: Invalid fail criteria %s. Supported criteria are: %v", constants.CI_FAILURE, failCriteria, constants.FAIL_CRITERIA_SEVERITIES)
				os.Exit(1)
			}

		case constants.DIGGITY:
			if len(failCriteria) > 0 {
				log.Fatalf("%v: Fail criteria is not supported for diggity analyzer", constants.CI_FAILURE)
				os.Exit(1)
			}
		default:
			log.Fatalf("%v: No analyzer type %s. Use --analyzer flag to provide an analyzer type. Choose: %v", constants.CI_FAILURE, analyzer, constants.SUPPORTED_ANALYZERS)
			os.Exit(1)
		}
	} else {
		log.Fatalf("%v: No analyzer type %s. Use --analyzer flag to provide an analyzer type. Choose: %v", constants.CI_FAILURE, analyzer, constants.SUPPORTED_ANALYZERS)
		os.Exit(1)
	}

	// # SCAN TYPE FLAG
	if len(scanType) == 0 && !helper.Contains(constants.SUPPORTED_SCAN_TYPES[:], scanType) {
		log.Fatalf("%v: Invalid scan type %s. Supported types are: %v", constants.CI_FAILURE, scanType, constants.SUPPORTED_SCAN_TYPES)
		os.Exit(1)
	}

	// # INPUT FLAG
	if len(input) == 0 {
		_ = c.Help()
		os.Exit(0)
	}

	// # API TAGS: TOKEN & PLUGIN TYPE FLAGS
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
	// # ENVIRONMENT TYPE FLAG
	if len(environmentType) == 0 {
		log.Fatalf("%v: No Environment Type provided. Use --environment-type flag to provide an environment type.", constants.CI_FAILURE)
		os.Exit(1)
	}

	if len(environmentType) > 0 && !helper.Contains(constants.SUPPORTED_ENVIRONMENT_TYPE[:], environmentType) {
		log.Fatalf("%v: Invalid environment type %s. Supported environment types are: %v", constants.CI_FAILURE, environmentType, constants.SUPPORTED_ENVIRONMENT_TYPE)
		os.Exit(1)
	}

	// Validate Types
	envType := util.GetEnvironmentType(environmentType)
	analysisType := util.GetAnalysisType(analyzer)

	// # SET PARAMETERS
	parameters := types.Parameters{
		Analyzer:        analyzer,
		ScanType:        scanType,
		Input:           input,
		Token:           token,
		PluginType:      pluginType,
		FailCriteria:    failCriteria,
		SkipFail:        skipFail,
		ForceDbUpdate:   forceDbUpdate,
		EnvironmentType: environmentType,
		EnvType:         envType,
		AnalysisType:    analysisType,

		Diggity: diggity.Parameters{
			OutputFormat: diggity.JSON,
		},

		Jacked: jacked.Parameters{
			Format: jacked.JSON,
		},
	}

	oss.Run(parameters)

}
