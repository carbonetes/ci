package jacked

import (
	"github.com/CycloneDX/cyclonedx-go"
	"github.com/carbonetes/ci/internal/db"
	"github.com/carbonetes/jacked/pkg/analyzer"
)

func Analyze(bom *cyclonedx.BOM) {

	analyzer.AnalyzeCDX(bom)
}

func DBRun(forceDbUpdate bool) {

	db.DBCheck(false, forceDbUpdate)
	db.Load()
}
