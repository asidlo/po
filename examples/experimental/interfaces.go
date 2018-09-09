package main

import "fmt"

type series []entry
type entry interface{}

type strentry string
type intentry int

type dataframe struct {
	data map[string]series
}

func (df dataframe) Select(c ...string) dataframe {
	dfd := make(map[string]series, len(df.data))
	for _, col := range c {
		if _, ok := df.data[col]; ok {
			dfd[col] = df.data[col]
		}
	}
	return dataframe{data: dfd}
}

func main() {
	an := "name1"
	kn := "name2"
	names := []entry{an, kn}

	aa := 25
	ka := 24
	ages := []entry{aa, ka}

	df := dataframe{
		data: map[string]series{
			"Name": names,
			"Age":  ages,
		},
	}
	// for k, v := range df.data {
	// 	fmt.Println(k, v)
	// 	for i, e := range v {
	// 		fmt.Println(i, e)
	// 	}
	// }
	fmt.Println(df)
	dfn := df.Select("Name")
	fmt.Println(dfn)
}
