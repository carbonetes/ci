package diggity

import (
	"os"
	"time"

	"github.com/CycloneDX/cyclonedx-go"
	"github.com/carbonetes/ci/cmd/ci/oss/diggity/secrets"
	"github.com/carbonetes/ci/cmd/ci/oss/jacked"
	"github.com/carbonetes/ci/internal/log"
	"github.com/carbonetes/ci/internal/presenter"

	"github.com/carbonetes/ci/internal/constants"
	"github.com/carbonetes/ci/pkg/types"
	"github.com/carbonetes/diggity/pkg/cdx"
	"github.com/carbonetes/diggity/pkg/reader"
	diggity "github.com/carbonetes/diggity/pkg/types"
)

func Analyze(parameters types.Parameters) (*cyclonedx.BOM, []diggity.Secret) {

	var bom *cyclonedx.BOM
	// Start Duration
	start := time.Now()

	addr, err := diggity.NewAddress()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	cdx.New(addr)

	switch parameters.Diggity.ScanType {
	case 1: // IMAGE
		image, ref, err := reader.GetImage(parameters.Input, nil)
		if err != nil {
			log.Printf("%v: Error reading image '%s': %v", constants.CI_FAILURE, parameters.Input, err)
			os.Exit(1)
		}

		cdx.SetMetadataComponent(addr, cdx.SetImageMetadata(*image, *ref, parameters.Input))

		err = reader.ReadFiles(image, addr)
		if err != nil {
			log.Printf("%v: Error reading files from image '%s': %v", constants.CI_FAILURE, parameters.Input, err)
			os.Exit(1)
		}
	case 2: // TARBALL
		tarball, err := reader.ReadTarball(parameters.Input)
		if err != nil {
			log.Printf("%v: Error reading tarball '%s': %v", constants.CI_FAILURE, parameters.Input, err)
			os.Exit(1)
		}
		err = reader.ReadFiles(tarball, addr)
		if err != nil {
			log.Printf("%v: Error reading files from tarball '%s': %v", constants.CI_FAILURE, parameters.Input, err)
			os.Exit(1)
		}
	case 3: // FILESYSTEM
		err := reader.FilesystemScanHandler(parameters.Input, addr)
		if err != nil {
			log.Printf("%v: Error reading file system '%s': %v", constants.CI_FAILURE, parameters.Input, err)
			os.Exit(1)
		}
	default:
		log.Printf("%v: Unsupported scan type '%s'. Supported scan types are: %v", constants.CI_FAILURE, parameters.ScanType, constants.SUPPORTED_SCAN_TYPES)
		os.Exit(1)
	}

	// # SBOM: Diggity Analysis Result
	bom = cdx.Finalize(addr)

	// # Secrets: Diggity Secrets Analysis Result
	secrets := secrets.Analyze()

	// # Vulnerability: Jacked Analysis
	if parameters.Analyzer == constants.JACKED {
		jacked.DBRun(parameters.ForceDbUpdate)
		start = time.Now()
		jacked.Analyze(bom)
	}
	// End Duration
	elapsed := time.Since(start).Seconds()
	parameters.Duration = time.Now().Add(-time.Duration(elapsed * float64(time.Second)))

	// Display Output
	run := presenter.DisplayAnalysisOutput(parameters, elapsed, bom, secrets)
	presenter.DisplayAssesstmentOutput(run, parameters)
	return bom, secrets
}
