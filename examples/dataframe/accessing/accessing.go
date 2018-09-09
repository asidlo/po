package main

import (
	"fmt"

	"github.com/asidlo/po"
)

func main() {
	s1 := po.Series{"test", "1", "7.5"}
	s2 := po.Series{"example", "2", "6.9"}
	df := po.NewDataFrame([]po.Series{s1, s2}, []string{"label", "index", "weight"})

	// Get Column Names for DataFrame
	cols := df.Columns()
	fmt.Println(cols)

	// Accessing Columns from DataFrame
	col1 := df[cols[0]]
	fmt.Println(col1)
}
