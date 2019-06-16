package lib

import (
	"errors"
	"fmt"
)

type OperationCount struct {
	column string
}

func (this *OperationCount) Construct(source *Source, args []string) (error, []string) {
	if len(args) < 2 || args[1] == "" {
		return errors.New("--count operation requires a column name."), []string{}
	}
	this.column = args[1]
	return nil, args[2:]
}

func (this *OperationCount) Run(source *Source) error {
	count := len(source.Rows) - 1
	source.Rows = [][]string{
		{this.column},
		{fmt.Sprintf("%d", count)},
	}
	return nil
}
