package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hollowness-inside/SBVJ01-Reader/pkg/sbvj"
)

func main() {
	// Writing
	{
		file, err := os.Create("data/output.sbvj")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		options := sbvj.SBVJOptions{
			Name:      "TestFile",
			Versioned: true,
			Version:   1234,
		}

		wr, err := sbvj.NewWriter(file, &options)
		if err != nil {
			log.Fatal(err)
		}

		if err := wr.PackString("Hello World"); err != nil {
			log.Fatal(err)
		}

		wr.Flush()
	}

	// Reading
	{
		sbvj, err := sbvj.ReadFile("data/output.sbvj")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Name:", sbvj.Name)
		fmt.Printf("Versioned (%t) = %d\n", sbvj.Versioned, sbvj.Version)
		fmt.Println(sbvj.Content.Value.(string))
	}

	// Output:
	// Name: TestFile
	// Versioned (true) = 1234
	// Hello World
}
