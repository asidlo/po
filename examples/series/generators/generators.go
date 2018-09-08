package main

import (
	"fmt"

	"github.com/asidlo/po"
)

func main() {
	// IntGenerator (start, end, step, exlude)
	// Tip: Good when used as a parameter for taking a subset
	is := po.IntGenerator(0, 10, 2, []int{4})
	fmt.Println(is)
}
