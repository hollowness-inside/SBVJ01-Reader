package main

import (
	"fmt"
	"joshua/green/starbound/sbvj01"
)

func main() {
	data, _ := sbvj01.ReadSBVJ01File("file.player")

	fmt.Println("Name:", data.Name)
	fmt.Printf("Versioned (%t) = %d", data.Versioned, data.Version)
	fmt.Println(data.Value)
}
