package main

import (
	"encoding/csv"
	"fmt"
	"github.com/evantbyrne/csvutil/lib"
	"os"
)

func main() {
	args := os.Args[1:]
	source := lib.ArgList(args)
	if source == nil {
		return
	}

	source.Run()

	w := csv.NewWriter(os.Stdout)
	w.WriteAll(source.Rows)
	w.Flush()
	if err := w.Error(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
