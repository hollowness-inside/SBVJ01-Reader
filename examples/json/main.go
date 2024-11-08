package main

import (
	"encoding/json"
	"os"

	"github.com/hollowness-inside/SBVJ01-Reader/pkg/sbvj"
)

func main() {
	sbvj, err := sbvj.ReadFile("data/file.player")
	if err != nil {
		panic(err)
	}

	content := sbvj.Content

	jsoned, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	output, err := os.Create("data/output.json")
	if err != nil {
		panic(err)
	}

	_, err = output.Write(jsoned)
	if err != nil {
		panic(err)
	}
}
