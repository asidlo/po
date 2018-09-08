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
}
