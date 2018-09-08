package main

import (
	"fmt"

	"github.com/asidlo/po"
)

func main() {
	s1 := po.Series{"test", "1", "7.5"}
	s2 := po.Series{"example", "2", "6.9"}
	df := po.NewDataFrame([]po.Series{s1, s2}, []string{"label", "index", "weight"})

	// Pick a Subset of a DataFrame via Row Index
	b := df.Pick([]int{1, 2, 3, 4}...)
	fmt.Println(b)

	// Select a Subset of a DataFrame via range
	dfs := df.Subset(0, 5, 2, nil)
	fmt.Println(dfs)
}
