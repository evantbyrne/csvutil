package lib

import (
	"errors"
)

type OperationWhere struct {
	comparison *Comparison
}

func (this *OperationWhere) Construct(args []string) (error, []string) {
	if len(args) < 4 {
		return errors.New("--where operation requires three arguments."), []string{}
	}
	err, comparison, remainingArgs := ConstructComparison("--where", args[1], args[2], args[3:])
	if err != nil {
		return err, []string{}
	}
	this.comparison = comparison
	return nil, remainingArgs
}

func (this *OperationWhere) Run(source *Source) error {
	var rows [][]string
	for _, row := range source.Rows {
		if this.comparison.Match(source, row) {
			rows = append(rows, row)
		}
	}
	source.Rows = rows
	return nil
}
