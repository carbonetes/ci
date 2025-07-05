package presenter

import (
	"strings"

	"github.com/CycloneDX/cyclonedx-go"
	"github.com/carbonetes/ci/cmd/ci/ui/table"
	"github.com/carbonetes/ci/internal/constants"
	"github.com/carbonetes/ci/internal/log"
	"github.com/carbonetes/ci/pkg/types"
)

func DisplayInput(parameters types.Parameters) {
	log.Println("========================================")
	log.Println("         Analysis Started")
	log.Println("========================================")
	log.Printf("  Analyzer    : %s", parameters.Analyzer)
	log.Printf("  Input       : %s", parameters.Input)
	log.Printf("  Scan Type   : %s", parameters.ScanType)
	log.Printf("  Plugin Type : %s", parameters.PluginType)
	if parameters.Analyzer == constants.JACKED {
		log.Printf("  Fail Criteria: %s", parameters.FailCriteria)
	}
	log.Println("========================================")
}
func DisplayOutput(parameters types.Parameters, duration float64, bom *cyclonedx.BOM) {
	switch parameters.Analyzer {
	case constants.JACKED:
		if bom == nil || bom.Vulnerabilities == nil || bom.Components == nil {
			log.Printf("No vulnerabilities found in BOM")
			log.Printf("Analysis completed in %.3f seconds", duration)
			return
		}

		tbl := table.NewTable()
		tbl.SetHeaders("Component", "Version", "CVE", "Severity", "Recommendation")

		// Build a map of BOMRef to component name:version
		componentsMap := make(map[string]string)
		for _, c := range *bom.Components {
			componentsMap[c.BOMRef] = c.Name + ":" + c.Version
		}

		for _, v := range *bom.Vulnerabilities {
			component, ok := componentsMap[v.BOMRef]
			if !ok {
				log.Printf("Component not found for vulnerability: %s", v.BOMRef)
				continue
			}
			parts := strings.Split(component, ":")
			name := parts[0]
			version := ""
			if len(parts) > 2 {
				version = strings.Join(parts[1:], ":")
			} else if len(parts) == 2 {
				version = parts[1]
			}

			severity := "UNKNOWN"
			if v.Ratings != nil && len(*v.Ratings) > 0 {
				for _, r := range *v.Ratings {
					if r.Severity != "" {
						severity = string(r.Severity)
						break
					}
				}
			}

			tbl.AddRow(
				name,
				version,
				v.ID,
				severity,
				v.Recommendation,
			)
		}

		tbl.Print()
		log.Printf("Analysis completed in %.3f seconds", duration)
	case constants.DIGGITY:
		if bom == nil || bom.Components == nil {
			log.Printf("No components found in BOM")
			log.Printf("Analysis completed in %.3f seconds", duration)
			return
		}

		tbl := table.NewTable()
		tbl.SetHeaders("Package Name", "Type", "Version")

		for _, c := range *bom.Components {
			componentType := ""
			if c.Properties != nil {
				for _, p := range *c.Properties {
					if p.Name == "diggity:package:type" && p.Value != "" {
						componentType = p.Value
						break
					}
				}
			}
			tbl.AddRow(
				c.Name,
				componentType,
				c.Version,
			)
		}

		tbl.Print()
		log.Printf("Analysis completed in %.3f seconds", duration)
	}
}
