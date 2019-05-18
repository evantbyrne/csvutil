package lib

import (
	"errors"
)

type OperationExcept struct{}

func (this *OperationExcept) Construct(source *Source, args []string) (error, []string) {
	return nil, args[1:]
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
