package main

import (
	"fmt"

	"github.com/asidlo/po"
)

func main() {
	rdf := po.GenerateDataFrame(20)
	cols := rdf.Cols()
	c := cols[0:5]
	df := rdf.Select(c...)
	fmt.Println(df)
}
