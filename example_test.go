package sbvj

import (
	"fmt"
	"log"
	"os"
)

func ExampleReadFile() {
	sbvj, err := ReadFile("testdata/file.player")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Name:", sbvj.Name)
	fmt.Printf("Versioned (%t) = %d\n", sbvj.Versioned, sbvj.Version)

	data := sbvj.Value.Value.(SBVJMap)

	movCont := data.Get("movementController").Value.(SBVJMap)
	facDir := movCont.Get("facingDirection").Value.(string)

	fmt.Println("Player facing direction:", facDir)

	// Output:
	// Name: PlayerEntity
	// Versioned (true) = 31
	// Player facing direction: right
}

func ExampleWrite() {
	file, err := os.Create("testdata/output.sbvj")
	if err != nil {
		log.Fatal(err)
	}

	wr := NewWriter(file)
	defer wr.Flush()

	if err := wr.PackString("Hello World"); err != nil {
		log.Fatal(err)
	}
}
