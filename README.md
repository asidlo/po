# Po - Pandas for Go

Data science library inspired by the python pandas library and named after the Po from [Kung Fu Panda](https://en.wikipedia.org/wiki/Kung_Fu_Panda). Po allows the user to perform data munging/wrangling techniques by modeling the data as a DataFrame. Much like the dataframe from the Pandas library, the DataFrame construct in Po is comprised of 0..n Series. However, unlike the Pandas implementation, all values in the Po Series are instantiated as strings. The rationale behind this design was to minimize the end user learning curve by mimicing the panda's api as close as possible. Without the aid of overly complicated generics and under the beautifully strict enforcement of types, binding data to strings using go's native slice and map types seemed like the best approach to achieve that emulation. So in accordance with that reasoning, Series in Po are types of slice[string] and DataFrames are types of map[string]string.

Despite this seemingly limited binding, Po still allows users to perform analytical operations on DataFrames and Series. This is achieved by converting the Series in question to other defined types in Po, namely IntSlice and FloatSlice. Series, IntSlice, and FloatSlice all have methods allowing the user to convert from one to the other in order to perform operations on a Series that would apply to the nature of the respective datatype. Future iterations will expand on the these defined types to include types such as boolean, time.Time, map, interface, rune, and possible others.

## Install

```bash
go get github.com/asidlo/po
```

## Examples

DataFrame Example

```go
s1 := po.Series{"1", "2", "3"}
s2 := po.Series{"8", "7", "6"}
df3 := po.NewDataFrame([]po.Series{s1, s2}, []string{"a", "b", "c"})
x, y = df3.Dims()
fmt.Printf("Normal DataFrame example: (%d x %d)\n%s\n", x, y, df3)
```

Series Example

```go
s1 := po.Series{
  "This", "Is", "1", "Example", "Of", "A", "Series",
}
fmt.Println(s1)
```

More examples can be found in the [examples](examples) directory.

## Alternatives

- [github.com/kniren/gota](https://github.com/kniren/gota/)