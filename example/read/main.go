package main

import (
	"fmt"
	"log"

	"github.com/hollowness-inside/SBVJ01-Reader/pkg/sbvj"
	"github.com/hollowness-inside/SBVJ01-Reader/pkg/types"
)

func main() {
	sbvj, err := sbvj.ReadFile("data/file.player")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Name:", sbvj.Name)
	fmt.Printf("Versioned (%t) = %d\n", sbvj.Versioned, sbvj.Version)

	data := sbvj.Content.Value.(types.SBVJMap)

	movCont := data.Get("movementController").Value.(types.SBVJMap)
	facDir := movCont.Get("facingDirection").Value.(string)

	fmt.Println("Player facing direction:", facDir)

	// Output:
	// Name: PlayerEntity
	// Versioned (true) = 31
	// Player facing direction: right
}
