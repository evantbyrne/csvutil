package lib

import (
	"errors"
	"strings"
)

type OperationDistinct struct {
	columns string
}

func (this *OperationDistinct) Construct(source *Source, args []string) (error, []string) {
	if len(args) < 2 || args[1] == "" {
		return errors.New("--distinct operation requires a comma-separated list of column names to compare. Use \"*\" to compare all."), []string{}
	}
	this.columns = args[1]
	return nil, args[2:]
}

func (this *OperationDistinct) Run(source *Source) error {
	var columnIndexes []int
	if this.columns == "*" {
		for i := range source.Rows[0] {
			columnIndexes = append(columnIndexes, i)
		}
	} else {
		for _, key := range strings.Split(this.columns, ",") {
			err, columnIndex := source.ColumnIndex(key)
			if err != nil {
				return err
			}
			columnIndexes = append(columnIndexes, columnIndex)
		}
	}

	var rows [][]string
	for _, row := range source.Rows {
		exclude := false
		for _, rowIncluded := range rows {
			allMatch := true
			for _, columnIndex := range columnIndexes {
				if row[columnIndex] != rowIncluded[columnIndex] {
					allMatch = false
					break
				}
			}
			if allMatch {
				exclude = true
				break
			}
		}
		if !exclude {
			rows = append(rows, row)
		}
	}
	source.Rows = rows
	return nil
}
