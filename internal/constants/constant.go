package constants

const (
	CI_FAILURE = "#CI-FAILED: "
	CI_SUCCESS = "#CI-SUCCESS: "
	CI_WARNING = "#CI-WARNING: "
	DIGGITY    = "diggity"
	JACKED     = "jacked"

	// SCAN TYPES
	IMAGE       = "image"
	FILE_SYSTEM = "filesystem"
	TARBALL     = "tarball"
)

// SUPPORTED CI/CD PLUGINS
var SUPPORTED_CICD_PLUGINS = [...]string{"jenkins", "azure", "bitbucket"}

// FAIL CRITERIA SEVERITIES
var FAIL_CRITERIA_SEVERITIES = [...]string{"critical", "high", "medium", "low", "negligible", "unknown"}

// SUPPORTED SCAN TYPES
var SUPPORTED_SCAN_TYPES = [...]string{"image", "filesystem", "tarball"}

// SUPPORTED ANALYZERS
var SUPPORTED_ANALYZERS = [...]string{JACKED, DIGGITY}
