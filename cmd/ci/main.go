package main

import (
	"github.com/carbonetes/ci/internal/log"

	"github.com/carbonetes/ci/cmd/ci/command"
)

func main() {
	if err := command.Run(); err != nil {
		log.Fatal(err)
	}
}
