package main

import (
	"encoding/csv"
	"github.com/estenrye/csvsplitter/csvsplitter"
	"flag"
	"fmt"
	"os"
)

func main() {
	splitColumn := flag.Int("splitColumn", 1, "Index of the columnt to split the file on. Default of 1 is the leftmost column.")
	hasHeaderRow := flag.Bool("hasHeaderRow", false, "Enables splitting data with a header row.")
	flag.Parse()

	if len(flag.Args()) < 1	{
		fmt.Print("No filename provided.")
		return
	}

	for _,s := range(flag.Args()) {
		f, err := os.Open(s)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		r := csv.NewReader(f)

		fmt.Println("Reading file: " + s)
		rows, err := r.ReadAll()
		if err != nil {
			fmt.Println(err)
			return
		}

		files := csvsplitter.SplitRows(rows, *hasHeaderRow, *splitColumn - 1)

		for _, data := range(files) {
			filename := data.Key + ".csv"
			out, err := os.Create(filename)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer out.Close()

			w := csv.NewWriter(out)

			if data.HeaderRow != nil {
				w.Write(data.HeaderRow)
			}
			w.WriteAll(data.Rows)
		}
	}
}
