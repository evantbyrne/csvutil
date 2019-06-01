package main

import (
	"encoding/csv"
	"fmt"
	"github.com/evantbyrne/csvutil/lib"
	"os"
)

func main() {
	args := os.Args[1:]
	err, source := lib.ArgList(args, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if source == nil {
		return
	}

	if err := source.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	w := csv.NewWriter(os.Stdout)
	w.WriteAll(source.Rows)
	w.Flush()
	if err := w.Error(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
