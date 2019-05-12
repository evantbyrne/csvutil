package lib

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Source struct {
	Operations []Operation
	Previous   *Source
	Path       string
	Rows       [][]string
}

func (this *Source) ColumnIndex(key string) int {
	for i, value := range this.Rows[0] {
		if key == value {
			return i
		}
	}

	fmt.Printf("Invalid column '%s' for '%s'.\n", key, this.Path)
	os.Exit(1)
	return 0
}

func (this *Source) ContainsRow(row []string) bool {
	for _, b := range this.Rows {
		if ArraysEqual(row, b) {
			return true
		}
	}

	return false
}

func (this *Source) ReadAll() {
	fh, err := os.Open(this.Path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fh.Close()

	rows, err := csv.NewReader(fh).ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(rows) == 0 {
		fmt.Printf("Empty source '%s'.\n", this.Path)
		os.Exit(1)
	}

	this.Rows = rows
}

func (this *Source) Run() {
	if this.Previous != nil {
		this.Previous.Run()
	}

	this.ReadAll()
	for _, operation := range this.Operations {
		if err := operation.Run(this); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
