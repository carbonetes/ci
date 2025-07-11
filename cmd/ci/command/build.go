package command

import (
	"os"

	"github.com/carbonetes/ci/cmd/ci/build"
	"github.com/carbonetes/ci/internal/helper"
	"github.com/carbonetes/ci/internal/log"
	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Display the version of Carbonetes CI",
		Run:   versionRun,
	}
	format string = "text"
)

func init() {
	root.AddCommand(versionCmd)
	versionCmd.Flags().StringVarP(&format, "format", "f", "text", "Output format (text or json)")
}

func versionRun(c *cobra.Command, _ []string) {
	info := build.GetBuild()
	if format == "json" {
		output, err := helper.ToJSON(info)
		if err != nil {
			log.Printf("Error marshalling version info to JSON: %v", err)
		}
		log.Infof("%v", string(output))
		os.Exit(0)
	} else if format == "text" {
		log.Infof("Application\t: %v", info.Application)
		log.Infof("Version\t\t: %v", info.Version)
		log.Infof("Build Date\t: %v", info.BuildDate)
		log.Infof("Git Commit\t: %v", info.GitCommit)
		log.Infof("Git Description\t: %v", info.GitDesc)
		log.Infof("Go Version\t: %v", info.GoVersion)
		log.Infof("Compiler\t: %v", info.Compiler)
		log.Infof("Platform\t: %v", info.Platform)
		os.Exit(0)
	} else {
		log.Fatal("Invalid output format")
	}
}
