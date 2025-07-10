package jacked

import (
	"github.com/CycloneDX/cyclonedx-go"
	"github.com/carbonetes/jacked/pkg/analyzer"
	jackedDB "github.com/carbonetes/jacked/pkg/db"
)

func Analyze(bom *cyclonedx.BOM) {

	analyzer.Analyze(bom)
}

func DBRun(forceDbUpdate bool) {

	jackedDB.Check(forceDbUpdate)
	jackedDB.Load()
}
