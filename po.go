// Package po implements a Pandas-like library for go.
// Po provides Series and DataFrame datastructures for data munging and preparation.
// Inspired by https://github.com/pandas-dev/pandas and https://github.com/kniren/gota
package po

import (
	"encoding/csv"
	"io"
	"math/rand"
	"sort"
	"strconv"
	"strings"

	"github.com/Pallinder/go-randomdata"

	"github.com/olekukonko/tablewriter"
)

const (
	// HeadSize is the default return length of the Head()
	// for both a Series and a DataFrame
	HeadSize = 5

	// RandStringLen is the default length for the randomly generated
	// column names, if and when they are needed.
	RandStringLen = 5
)

// Series is a generic datastructure that contains a slice of strings.
// Strings were chosen as the type of choice since I wanted to
// make the api simple, easy to use, and as close to the pandas
// api as possible so that the learning curve would be small.
// This allows the user to input any type of data they want into a Series,
// (so long as it is surrounded by ""). There are casting operations that can
// be performed on a Series to perform different mathmatical operations which
// require non string types.
type Series []string

// DataFrame is a map datastructure containing Series values.
// It is intented to represent a generic table where the keys
// correspond to the individual column names and the rows correspond
// to the original input series. Column names can be generated on
// DataFrame instantiation via literal construct, or via the NewDataFrame()
type DataFrame map[string]Series

// Cols returns all of the column names sorted for maintaining order.
// The sorting is done using the sort.String() method from the std lib.
func (df DataFrame) Cols() []string {
	var keys []string
	for k := range df {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// Head returns the first i entries in a series.
// If a slice of ints is passed, only the first entry is used.
// I used the varargs operator to allow for optional entry.
// In the case where no i is passed, then default to returning
// the HeadSize or the len(Series) whichever is smaller.
// If a neg value is passed then the abs value is used.
func (s Series) Head(i ...int) Series {
	limit := len(s)
	if len(i) > 0 {
		// Check for negative
		i0 := i[0]
		if i0 < 0 {
			i0 = -i0
		}
		// Use passed head size
		if i0 < limit {
			limit = i0
		}
	} else {
		// If no param
		if HeadSize < limit {
			limit = HeadSize
		}
	}
	return s[0:limit]
}

// String returns the string representation of the Series.
// It uses the olekukonko/tablewriter library to render the table.
func (s Series) String() string {
	var sb strings.Builder
	table := tablewriter.NewWriter(&sb)
	table.SetBorder(false)
	table.SetColumnSeparator("")
	for i, v := range s {
		si := strconv.Itoa(i)
		table.Append([]string{si, v})
	}
	table.Render()
	return sb.String()
}

// Subset returns a subset of the original DataFrame. It grabs entries from
// each column by their indices, starting from the start int to the end int
// using the specified step size and excluding any indices specified.
func (s Series) Subset(start int, end int, step int, exclude []int) Series {
	indices := IntGenerator(start, end, step, exclude)
	return s.Pick(indices...)
}

// Pick returns a subset DataFrame comprised only of rows indices specified.
func (s Series) Pick(i ...int) Series {
	l := len(s)
	var series Series
	for i := range i {
		if i < l {
			series = append(series, s[i])
		}
	}
	return series
}

// Transpose returns a transposed DataFrame of the original DataFrame.
// The transposed column names become a string of the former row index.
func (df DataFrame) Transpose() DataFrame {
	r, c := df.Dims()
	dft := make(DataFrame, r)
	cols := df.Cols()
	for i := 0; i < r; i++ {
		index := strconv.Itoa(i)
		col := make(Series, c)
		dft[index] = col
		for j, v := range cols {
			dft[index][j] = df[v][i]
		}
	}
	return dft
}

// String returns the string representation of the DataFrame.
// Columns are ordered via sort.Strings() method.
// It uses the olekukonko/tablewriter library to render the table.
func (df DataFrame) String() string {
	var sb strings.Builder
	table := tablewriter.NewWriter(&sb)
	table.SetBorder(false)
	table.SetCenterSeparator(" ")
	table.SetAutoFormatHeaders(false)
	table.SetRowLine(false)
	table.SetColumnSeparator(" ")
	table.SetHeaderLine(true)
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	cols := df.Cols()
	table.SetHeader(cols)
	dft := df.Transpose()
	for _, v := range dft {
		table.Append(v)
	}
	table.Render()
	return sb.String()
}

// Head returns the first i entries for each column in a DataFrame.
// If a slice of ints is passed, only the first entry is used.
// I used the varargs operator to allow for optional entry.
// In the case where no i is passed, then default to returning
// the HeadSize or the len DataFrame rows whichever is smaller.
// If a neg value is passed then the abs value is used.
func (df DataFrame) Head(i ...int) DataFrame {
	limit, _ := df.Dims()
	if len(i) > 0 {
		// Check for negative
		i0 := i[0]
		if i0 < 0 {
			i0 = -i0
		}
		// Use passed head size
		if i0 < limit {
			limit = i0
		}
	} else {
		// If no param
		if HeadSize < limit {
			limit = HeadSize
		}
	}
	// TODO Take a subset of the DataFrame
	return df.Subset(0, limit, 1, nil)
}

// IntGenerator generates a slice of integers that can then be used
// to apply different functions on DataFrames or Series, such as
// Subset() or Pick().
func IntGenerator(start int, end int, step int, exclude []int) []int {
	im := make(map[int]int, len(exclude))
	for _, v := range exclude {
		im[v] = v
	}
	var ints []int
	for i := 0; i < end; i += step {
		_, ok := im[i]
		if !ok {
			ints = append(ints, i)
		}
	}
	return ints
}

// Subset returns a subset of the original DataFrame. It grabs entries from
// each column by their indices, starting from the start int to the end int
// using the specified step size and excluding any indices specified.
func (df DataFrame) Subset(start int, end int, step int, exclude []int) DataFrame {
	indices := IntGenerator(start, end, step, exclude)
	return df.Pick(indices...)
}

// Pick returns a subset DataFrame comprised only of rows indices specified.
func (df DataFrame) Pick(i ...int) DataFrame {
	sdf := make(DataFrame, len(df))
	for k, v := range df {
		l := len(v)
		var s Series
		for _, index := range i {
			if index < l {
				s = append(s, v[index])
			}
		}
		sdf[k] = s
	}
	return sdf
}

// Dims returns the number of rows, number of columns in a DataFrame
func (df DataFrame) Dims() (int, int) {
	cols := df.Cols()
	if len(cols) > 0 {
		rowLen := len(df[cols[0]])
		colLen := len(cols)
		return rowLen, colLen
	}
	return 0, 0
}

// Shape returns the number of rows, number of columns in a DataFrame.
// Same as po.DataFrame.Dims()
func (df DataFrame) Shape() (int, int) {
	return df.Dims()
}

// Select returns a subset of the original DataFrame with only the given
// column names n represented.
func (df DataFrame) Select(c ...string) DataFrame {
	rdf := make(DataFrame, len(c))
	for _, col := range c {
		rdf[col] = df[col]
	}
	return rdf
}

// NewSeries is a variadic function that returns a Series comprised of the provided strings
func NewSeries(s ...string) Series {
	series := make(Series, len(s))
	for i, v := range s {
		series[i] = v
	}
	return series
}

// NewDataFrame returns a new DataFrame object with rows corresponding
// to provided ss and column names corresponding to provided cols.
// If no ss is provided, then an empty dataframe will be created using
// any provided column names. If the number of cols provided < len(ss)
// for any given ss, then the column names will be auto generated
// for the remaining entries. If the len(cols) > len(ss) for any given ss,
// then the ss will be extended with empty string values for each remaining col.
func NewDataFrame(ss []Series, cols []string) DataFrame {
	if len(ss) == 0 {
		df := make(DataFrame, len(cols))
		if len(cols) > 0 {
			for _, k := range cols {
				df[k] = Series{}
			}
		}
		return df
	}
	return createDataFrame(ss, cols)
}

// createDataFrame returns a new DataFrame supplementing each Series with
// additional empty string entries to make each Series of equal length.
// It also adds additional column names if necessary to match max series length.
func createDataFrame(ss []Series, cols []string) DataFrame {
	df := make(DataFrame, len(cols))

	// Find max length
	maxSeriesLen := 0
	for _, s := range ss {
		l := len(s)
		if maxSeriesLen < l {
			maxSeriesLen = l
		}
	}

	// Grow individual series accordingly
	for i, s := range ss {
		lDiff := maxSeriesLen - len(s)
		for j := 0; j < lDiff; j++ {
			ss[i] = append(ss[i], "")
		}
	}

	// Add columns if maxSeriesLen > len(cols)
	cDiff := maxSeriesLen - len(cols)
	if cDiff > 0 {
		for i := 0; i < cDiff; i++ {
			s := genRandStr(RandStringLen)
			cols = append(cols, s)
		}
	}

	// Add entries to series if len(cols) > maxSeriesLen
	if cDiff < 0 {
		cDiff = -cDiff
		for ix := range ss {
			for i := 0; i < cDiff; i++ {
				ss[ix] = append(ss[ix], "")
			}
		}
	}

	// Create DataFrame since now len(col) == maxLenSeries
	for c := 0; c < len(cols); c++ {
		col := cols[c]
		s := make(Series, len(ss))
		df[col] = s
		for r := 0; r < len(ss); r++ {
			df[col][r] = ss[r][c]
		}
	}
	return df
}

// genRanStr generates a random string with length n
// Shoutout to: https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func genRandStr(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n+1)
	// Prepend with 'z' for sorting, note 'Z' comes before 'a' in strings.sort
	b[0] = 'z'
	for i := 1; i < len(b); i++ {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// GenerateDataFrame generates a DataFrame with size n of randomized user profile data.
// The data is generated using the github.com/Pallinder/go-randomdata package
func GenerateDataFrame(n int) DataFrame {
	genders := make([]string, n)
	emails := make([]string, n)
	dobs := make([]string, n)
	registerds := make([]string, n)
	phones := make([]string, n)
	cells := make([]string, n)
	usernames := make([]string, n)
	firstnames := make([]string, n)
	lastnames := make([]string, n)

	for i := 0; i < n; i++ {
		profile := randomdata.GenerateProfile(randomdata.RandomGender)
		genders[i] = profile.Gender
		emails[i] = profile.Email
		dobs[i] = profile.Dob
		registerds[i] = profile.Registered
		phones[i] = profile.Phone
		cells[i] = profile.Cell
		usernames[i] = profile.Login.Username
		firstnames[i] = profile.Name.First
		lastnames[i] = profile.Name.Last
	}
	return DataFrame{
		"Gender":       genders,
		"Email":        emails,
		"DOB":          dobs,
		"RegisterDate": registerds,
		"PhoneNumber":  phones,
		"CellNumber":   cells,
		"UserName":     usernames,
		"FirstName":    firstnames,
		"LastName":     lastnames,
	}
}

// ReadCsv reads in csv data and returns a DataFrame with randomly generated
// column names for each input column.
func ReadCsv(r io.Reader) (DataFrame, error) {
	cr := csv.NewReader(r)
	records, err := cr.ReadAll()
	if err != nil {
		return nil, err
	}
	ss := make([]Series, len(records))
	for i, row := range records {
		s := NewSeries(row...)
		ss[i] = s
	}
	return NewDataFrame(ss, nil), nil
}

// WriteCsv writes a dataframe to the given io.Writer in csv format.
func WriteCsv(w io.Writer, df DataFrame) error {
	dft := df.Transpose()
	records := make([][]string, len(dft))
	for i, c := range dft.Cols() {
		records[i] = dft[c]
	}
	cw := csv.NewWriter(w)
	return cw.WriteAll(records)
}
