package lib

import (
	"errors"
	"fmt"
	"strings"
)

type OperationExcept struct {
	columns string
}

func (this *OperationExcept) Construct(source *Source, args []string) (error, []string) {
	if len(args) < 2 || args[1] == "" {
		return errors.New("--except operation requires a comma-separated list of column names to compare. Use \"*\" to compare all."), []string{}
	}
	this.columns = args[1]
	return nil, args[2:]
}

func (this *OperationExcept) Run(source *Source) error {
	if source.Previous == nil {
		return errors.New("--except operation cannot be run on first source.")
	}

	if !ArraysEqual(source.Previous.Rows[0], source.Rows[0]) {
		columns := strings.Join(source.Rows[0], ",")
		columnsPrevious := strings.Join(source.Previous.Rows[0], ",")
		return fmt.Errorf("--except operation requires both sources to have matching columns. Given sources have columns '%s' and '%s'. Consider using --select operation to resolve.", columnsPrevious, columns)
	}

	var columnIndexes []int
	if this.columns == "*" {
		for i, _ := range source.Rows[0] {
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
	for _, previousRow := range source.Previous.Rows {
		exclude := false
		for _, row := range source.Rows {
			allMatch := true
			for _, columnIndex := range columnIndexes {
				if previousRow[columnIndex] != row[columnIndex] {
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
			rows = append(rows, previousRow)
		}
	}
	source.Rows = rows
	return nil
}
