package main

import (
	"fmt"

	"github.com/asidlo/po"
)

func main() {
	s := po.Series{"1", "2", "3", "4", "5", "6", "7", "8"}
	fmt.Println(s.Head())
	fmt.Println(s.Head(3))
}
