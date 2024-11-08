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

	// File Options
	{
		opts := sbvj.Options

		fmt.Println("Name:", opts.Name)
		fmt.Printf("Versioned (%t) = %d\n", opts.Versioned, opts.Version)
	}

	// Content
	{
		content := sbvj.Content.Value.(types.SBVJMap)

		movController := content.Get("movementController").Value.(types.SBVJMap)
		facDir := movController.Get("facingDirection").Value.(string)

		fmt.Println("Movement Controller:", movController)
		fmt.Println("Player facing direction:", facDir)
	}
}
