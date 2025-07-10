package helper

import "strings"

// FailCriteriaSeverityMatchesFilter Checks for equal or higher severity than the fail criteria.
func FailCriteriaSeverityMatchesFilter(severity, failCriteria string) bool {
	severity = strings.ToLower(severity)
	failCriteria = strings.ToLower(failCriteria)
	severityOrder := []string{"critical", "high", "medium", "low", "negligible", "unknown"}
	failIdx := -1
	sevIdx := -1
	for i, s := range severityOrder {
		if s == failCriteria {
			failIdx = i
		}
		if s == severity {
			sevIdx = i
		}
	}
	return failIdx != -1 && sevIdx != -1 && sevIdx <= failIdx
}
