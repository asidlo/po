package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/asidlo/po"
)

func main() {
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
	// Reading in a csv
	r, err := po.ReadCsv(strings.NewReader(in))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)

	// Writing to a io.Writer in csv format
	werr := po.WriteCsv(os.Stdout, r)
	if err != nil {
		log.Fatal(werr)
	}
}
