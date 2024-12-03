package main

import (
	"testing"

	"github.com/hollowness-inside/SBVJ01-Reader/pkg/sbvj"
	"github.com/hollowness-inside/SBVJ01-Reader/pkg/types"
)

func BenchmarkSBVJRead(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		b.StartTimer()

		sbvj, err := sbvj.ReadFile("../data/file.player")
		if err != nil {
			b.Fatal(err)
		}

		_ = sbvj.Options.Name
		_ = sbvj.Options.Version

		content := sbvj.Content.Value.(types.SBVJMap)
		_ = content["movementController"]
	}
}
