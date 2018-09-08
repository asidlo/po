package main

import (
	"fmt"

	"github.com/asidlo/po"
)

func main() {
	// Create a Series via Slice literal
	s1 := po.Series{
		"This", "Is", "An", "Example", "Of", "A", "Series",
	}
	fmt.Println(s1)
}
