package lib

import (
	"errors"
)

type OperationExcept struct{}

func (this *OperationExcept) Construct(arg string) error {
	if arg != "" {
		return errors.New("--except operation does not take a value.")
	}
	return nil
}

func (this *OperationExcept) Run(source *Source) error {
	if source.Previous == nil {
		return errors.New("--except operation cannot be run on first source.")
	}

	var rows [][]string
	for _, previousRow := range source.Previous.Rows {
		if !source.ContainsRow(previousRow) {
			rows = append(rows, previousRow)
		}
	}
	source.Rows = rows
	return nil
}
