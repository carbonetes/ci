package util

const (
	BOM  = "bom"
	VULN = "vuln"
)

func AnalysisTypeSelector(analysisType int) string {
	var analysisTypeStr string

	switch analysisType {
	case 0:
		analysisTypeStr = BOM
	case 1:
		analysisTypeStr = VULN
	default:
		analysisTypeStr = BOM
	}

	return analysisTypeStr
}
