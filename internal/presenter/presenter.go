package presenter

import (
	"fmt"
	"strings"

	"github.com/CycloneDX/cyclonedx-go"
	"github.com/carbonetes/ci/cmd/ci/ui/table"
	"github.com/carbonetes/ci/internal/constants"
	"github.com/carbonetes/ci/internal/helper"
	"github.com/carbonetes/ci/internal/log"
	"github.com/carbonetes/ci/pkg/types"
	diggityTypes "github.com/carbonetes/diggity/pkg/types"
)

func DisplayInput(parameters types.Parameters) {
	log.Println("========================================")
	log.Println("         Analysis Started")
	log.Println("========================================")
	log.Printf("         Analyzer : %s", parameters.Analyzer)
	log.Printf("            Input : %s", parameters.Input)
	log.Printf("        Scan Type : %s", parameters.ScanType)
	log.Printf("      Plugin Type : %s", parameters.PluginType)
	log.Printf("        Skip Fail : %v", displaySkipFailonInput(parameters.SkipFail))
	if parameters.Analyzer == constants.JACKED {
		log.Printf("    Fail Criteria : %s", parameters.FailCriteria)
		if parameters.ForceDbUpdate {
			log.Printf("  Force DB Update : %v", parameters.ForceDbUpdate)
		}
	}
	log.Println("========================================")
	log.Println()
}

func DisplayAnalysisOutput(parameters types.Parameters, duration float64, bom *cyclonedx.BOM, secrets []diggityTypes.Secret) bool {
	switch parameters.Analyzer {
	case constants.JACKED:
		if bom == nil || bom.Components == nil || len(*bom.Vulnerabilities) == 0 {
			log.Println("========================================")
			log.Println("         Analysis Result")
			log.Println("========================================")
			log.Printf("       Vulnerabilities : %d", 0)
			log.Printf("Failure Criteria Found : %d", 0)
			log.Printf("              Duration : %.3f seconds", duration)
			log.Println("========================================")
			return true
		} else {

			tbl := table.NewTable()
			tbl.SetHeaders("Component", "CVE", "Version", "Recommendation", "Severity")

			// Build a map of BOMRef to component name:version
			componentsMap := make(map[string]string)
			for _, c := range *bom.Components {
				componentsMap[c.BOMRef] = c.Name + ":" + c.Version
			}

			// Count vulnerabilities that match the fail criteria
			failCriteriaCount := 0

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
				displaySeverity := severity
				if v.Ratings != nil && len(*v.Ratings) > 0 {
					for _, r := range *v.Ratings {
						if r.Severity != "" {
							severity = strings.ToLower(string(r.Severity))
							displaySeverity = severity
							if helper.FailCriteriaSeverityMatchesFilter(severity, parameters.FailCriteria) {
								displaySeverity = string(r.Severity) + "[!]"
								failCriteriaCount++
							}
							break
						}
					}
				}

				tbl.AddRow(
					name,
					v.ID,
					version,
					v.Recommendation,
					displaySeverity,
				)
			}

			tbl.Print()
			log.Println()
			log.Println("========================================")
			log.Println("         Analysis Result")
			log.Println("========================================")
			log.Printf("       Vulnerabilities : %d", len(*bom.Vulnerabilities))
			log.Printf("Failure Criteria Found : %d", failCriteriaCount)
			log.Printf("              Duration : %.3f seconds", duration)
			log.Println("========================================")
			return failCriteriaCount == 0
		}
	case constants.DIGGITY:
		if bom == nil || bom.Components == nil || len(*bom.Components) == 0 {
			log.Printf("No Packages Found")
			log.Printf("Duration: %.3f seconds", duration)
			return true
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
		log.Println()
		log.Println("========================================")
		log.Println("         Analysis Result")
		log.Println("========================================")
		if len(secrets) > 0 {
			log.Printf("      Secrets : %d [!]", len(secrets))
		} else {
			log.Printf("      Secrets : %d", len(secrets))
		}
		log.Printf("     Packages : %d", len(*bom.Components))
		log.Printf("     Duration : %.3f seconds", duration)
		log.Println("========================================")
		if len(secrets) > 0 {
			return false
		}
		return true
	}
	return true
}

func DisplayAssesstmentOutput(run bool, parameters types.Parameters) {
	if parameters.SkipFail {
		// Skip Fail == True Always return CI_SUCCESS
		DisplaySkipFail()
		run = true
	}
	log.Println()
	log.Println("========================================")
	if !run {
		log.Println(constants.CI_FAILURE, "Analysis completed with failure!")
	} else {
		log.Println(constants.CI_SUCCESS, "Analysis completed!")
	}
	log.Println("========================================")
}

func DisplaySkipFail() {
	log.Println()
	log.Printf("%v Skip fail is ENABLED!", constants.CI_WARNING)
	log.Println()
}

func displaySkipFailonInput(skipFail bool) string {
	if skipFail {
		return fmt.Sprintf("%t [WARNING]", skipFail)
	}
	return fmt.Sprintf("%t", skipFail)
}
