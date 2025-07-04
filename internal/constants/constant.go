package constants

const (
	CI_FAILURE = "#CI-FAILED: "
	CI_SUCCESS = "#CI-SUCCESS: "
	CI_WARNING = "#CI-WARNING: "
	DIGGITY    = "diggity"
	JACKED     = "jacked"
)

// SUPPORTED CI/CD PLUGINS
var SUPPORTED_CICD_PLUGINS = [...]string{"jenkins", "azure", "bitbucket"}

// FAIL CRITERIA SEVERITIES
var FAIL_CRITERIA_SEVERITIES = [...]string{"critical", "high", "medium", "low", "negligible", "unknown"}
