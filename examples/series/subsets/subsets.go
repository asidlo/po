package main

import (
	"fmt"

	"github.com/asidlo/po"
)

func main() {
	s := po.Series{"5", "sdkjwnefw", "wer", "jeu"}

	// Subsetting by index
	firstTwo := s.Pick(0, 1)
	fmt.Println(firstTwo)

	// Using Subset method to pick over range
	ss := s.Subset(0, len(s)-2, 1, nil)
	fmt.Println(ss)
}
