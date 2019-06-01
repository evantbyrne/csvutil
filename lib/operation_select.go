package lib

import (
	"errors"
	"strings"
)

type OperationSelect struct {
	columns []string
}

func (this *OperationSelect) Construct(source *Source, args []string) (error, []string) {
	if len(args) < 2 || args[1] == "" {
		return errors.New("--select operation requires a value."), []string{}
	}
	this.columns = strings.Split(args[1], ",")
	return nil, args[2:]
}

func (this *OperationSelect) Run(source *Source) error {
	var indices []int
	for _, key := range this.columns {
		err, columnIndex := source.ColumnIndex(key)
		if err != nil {
			return err
		}
		indices = append(indices, columnIndex)
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
