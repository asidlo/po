package main

import (
	"fmt"

	"github.com/asidlo/po"
)

func main() {
	s1 := po.Series{"test", "1", "7.5"}
	s2 := po.Series{"example", "2", "6.9"}
	df := po.NewDataFrame([]po.Series{s1, s2}, []string{"label", "index", "weight"})

	// Using Head(...int) to View First Few Entries of Series/DataFrame
	fmt.Println(df.Head())
	fmt.Println(df.Head(1))
}
