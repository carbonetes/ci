package types

import (
	"time"

	diggity "github.com/carbonetes/diggity/pkg/types"
	jacked "github.com/carbonetes/jacked/pkg/types"
)

type Parameters struct {
	// # ANALYZER PARAMETERS
	Analyzer string `json:"analyzer"`
	ScanType string `json:"scan_type"`
	Input    string `json:"input"`
	// # API PARAMETERS
	Token        string `json:"token"`
	PluginType   string `json:"plugin_type"`
	EnvType      int    `json:"env_type"`
	AnalysisType int    `json:"analysis_type"`
	// # GENERAL PARAMETERS
	FailCriteria    string    `json:"fail_criteria"`
	SkipFail        bool      `json:"skip_fail"`
	ForceDbUpdate   bool      `json:"force_db_update"`
	EnvironmentType string    `json:"environment_type"`
	Duration        time.Time `json:"duration"`

	// # DIGGITY & JACKED PARAMETERS
	Diggity diggity.Parameters `json:"diggity,omitempty"`
	Jacked  jacked.Parameters  `json:"jacked,omitempty"`
}
