package jacked

import (
	"github.com/CycloneDX/cyclonedx-go"
	"github.com/carbonetes/jacked/pkg/analyzer"
)

func Analyze(bom *cyclonedx.BOM) {

	// TODO: DB UPDATE FIRST BEFORE SCANNING
	analyzer.Analyze(bom)
}
