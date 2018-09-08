# TODO

links for pandas examples/docs:

- [Pandas Dataframes](https://pandas.pydata.org/pandas-docs/stable/api.html#dataframe)
- [Pandas Series](https://pandas.pydata.org/pandas-docs/stable/api.html#series)
- [Pandas Functions](https://paulovasconcellos.com.br/28-useful-pandas-functions-you-might-not-know-de42c59db085)
- [Panas Cheatsheet](https://www.dataquest.io/blog/pandas-cheat-sheet/)

Functions

- [ ] po.GenerateSeries(n int),GenerateFloat/IntSlice, mix and match
- [ ] ReadCsv()/WriteCsv()
- [ ] Tail()
- [ ] FFill()
- [ ] DropEmpty(all=true)/DropNull()
- [ ] FillEmpty(all=true, "")
- [ ] IsEmpty()/IsNull()
- [ ] po.DataFrame.Count() DataFrame [returns column and number of non null entries]
- [ ] po.DataFrame.DropCols([]string{"a", "b"}) [removing columns]
- [ ] po.DataFrame.DropRows([]int{0,1,2}) [removing rows]
- [ ] po.Series.Apply(func() po.Series)
- [ ] po.DataFrame.Rename(map[string]string)
- [ ] po.Series.Unique() po.Series
- [ ] po.Series.ToFloatSlice() po.FloatSlice
- [ ] po.Series.ToIntSlice() po.IntSlice
- [ ] po.FloatSlice.ToSeries() po.Series
- [ ] po.IntSlice.ToSeries() po.Series
- [ ] po.Float/IntSlice.Min()/Max()/Sum()/Mean()/Median()/
  - [ ] Min() and Max() should return (i int, value)
  - [ ] .Add(), .Sub(), .Mult(), .Div()
- [ ] sorting(asc/desc) for series/slices
- [ ] po.DataFrame.Join(df, on=, how=)
- [ ] po.DataFrame.concat([df1,df2]) add cols of 1 to 2 rows should be identical/filled
- [ ] DataFrame.append(df2) add rows in df2 to end of df1, cols will be created if not matching, and empty strings will be filled.
- [ ] po.DataFrame.Filter(...object) need some kind of struct for getting col name and comparator
- [ ] po.DataFrame.GroupBy(colNames []string creates grouped obj followed by agg
  - [ ] [link for groupby pandas examples](https://towardsdatascience.com/pandas-tips-and-tricks-33bcc8a40bb9)
- [ ] Describe()
- [ ] Info()
- [ ] Replace()
- [ ] Sample(n=, seed=)
- [ ] ReadSql()/WriteSql(xslx, "Sheet 1")
- [ ] ReadExcel()/WriteExcel(xslx, "Sheet 1")
- [x] po.GenerateDataFrame(n int) utilize other generators
- [x] po.DataFrame.Select(colNames []string) po.DataFrame

Extraneous

- [ ] Get Po Package working with GopherNotes
- [ ] Create Po Gopher Artwork, advertise on readme
- [ ] Write tests
- [ ] Make godocs fancy
- [ ] Introduce version of api
- [ ] Add changelog.md
- [ ] Order columns by user input instead of alphanumeric
- [ ] Adjust DataFrame String() method to handle large tables
- [ ] Add CONTRIBUTING.md