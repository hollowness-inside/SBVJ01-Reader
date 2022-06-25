package main

import (
	"fmt"
	"joshua/green/starbound/sbvj01"
)

func main() {
	data := sbvj01.ReadSBVJ01File("file.player")

	fmt.Println("Name:", data.Name)
	fmt.Printf("Versioned (%t) = %d\n", data.Versioned, data.Version)
	fmt.Printf("Type: %d\n", data.Value.Type)
	fmt.Println(data.Value)
}
