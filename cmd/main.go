package main

import (
	"fmt"

	"github.com/carbonetes/ci/util"
)

func main() {
	urls := []string{
		"http://localhost:3001/",
		"https://tent-api.carbonetes.com/",
		"https://prod.carbonetes.com/",
	}

	for _, url := range urls {
		enc, err := util.EncryptAESGCM(url)
		if err != nil {
			panic(err)
		}
		fmt.Println(enc)
	}
}
