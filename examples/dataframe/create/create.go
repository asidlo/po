package main

import (
	"fmt"

	"github.com/asidlo/po"
)

func main() {
	// Map literal instantiation (provided series will be represented as columns, instead of rows)
	// Note, needs to be nxn if using map literal and requires .Transform()
	// in order to have series in correct orientation
	df := po.DataFrame{
		"col1": po.Series{"a", "b"},
		"col2": po.Series{"1", "2"},
	}
	x, y := df.Dims()
	fmt.Printf("Map literal DataFrame Instantiation: (%d x %d)\n%s\n", x, y, df)

	// Empty DataFrame Instantiation Via NewDataFrame()
	df2 := po.NewDataFrame(nil, []string{"a", "b", "c"})
	x1, y1 := df2.Dims()
	fmt.Printf("Empty DataFrame Instantiation Via NewDataFrame(): (%d x %d)\n%s\n", x1, y1, df2)

	// Auto Grow Input Series to Match Input Column Length
	s1 := po.Series{"1"}
	s2 := po.Series{"8", "7", "6"}
	df3 := po.NewDataFrame([]po.Series{s1, s2}, []string{"a", "b", "c"})
	x2, y2 := df3.Dims()
	fmt.Printf("Auto Grow Input Series to match Input col len: (%d x %d)\n%s\n", x2, y2, df3)

	// Auto Generate Columns to Match Max Input Series Length
	s3 := po.Series{"1"}
	s4 := po.Series{"8", "7", "6"}
	df4 := po.NewDataFrame([]po.Series{s3, s4}, []string{"a"})
	x3, y3 := df4.Dims()
	fmt.Printf("Auto Gen Columns to match max input series: (%d x %d)\n%s\n", x3, y3, df4)
}
