package models

type Parameters struct {
	// # ANALYZER PARAMETERS
	Analyzer string `json:"analyzer"`
	ScanType string `json:"scan_type"`
	Input    string `json:"input"`
	// # API PARAMETERS
	Token      string `json:"token"`
	PluginType string `json:"plugin_type"`
	// # GENERAL PARAMETERS
	FailCriteria string `json:"fail_criteria"`
	SkipFail     bool   `json:"skip_fail"`
}
