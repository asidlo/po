package main

import (
	"fmt"

	"github.com/asidlo/po"
)

func main() {
	s1 := po.Series{"test", "1", "7.5"}
	s2 := po.Series{"example", "2", "6.9"}
	df := po.NewDataFrame([]po.Series{s1, s2}, []string{"label", "index", "weight"})
	fmt.Println(df)

	// Transposing DataFrame and Back Again
	fmt.Println("Taking transpose of a DataFrame:")
	dft := df.Transpose()
	fmt.Println(dft)

	// Renaming dataframe columns
	dft = dft.Rename(map[string]string{"0": "Col1", "1": "Col2"})
	fmt.Println(dft)

	// Removing columns from the dataframe
	dft = dft.DropColumns("Col2")
	fmt.Println(dft)
}
