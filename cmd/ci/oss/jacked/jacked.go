package jacked

import (
	"github.com/CycloneDX/cyclonedx-go"
	"github.com/carbonetes/jacked/pkg/analyzer"
)

func Analyze(bom *cyclonedx.BOM) {
	analyzer.Analyze(bom)
}
