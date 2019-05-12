package lib

import (
	"errors"
	"strings"
)

type OperationSelect struct {
	columns []string
}

func (this *OperationSelect) Construct(arg string) error {
	if arg == "" {
		return errors.New("--select operation requires a value.")
	}
	this.columns = strings.Split(arg, ",")
	return nil
}

func (this *OperationSelect) Run(source *Source) error {
	var indices []int
	for _, key := range this.columns {
		indices = append(indices, source.ColumnIndex(key))
	}

	for i, row := range source.Rows {
		var newRow []string
		for _, k := range indices {
			if k < len(row) {
				newRow = append(newRow, row[k])
			}
		}
		source.Rows[i] = newRow
	}
	return nil
}
