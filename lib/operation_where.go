package lib

import (
	"errors"
)

type OperationWhere struct {
	Comparisons []*Comparison
}

func (this *OperationWhere) Construct(source *Source, args []string) (error, []string) {
	if len(args) < 4 {
		return errors.New("--where operation requires three arguments."), []string{}
	}
	err, comparison, remainingArgs := ConstructComparison("--where", args[1], args[2], args[3:])
	if err != nil {
		return err, []string{}
	}
	this.Comparisons = append(this.Comparisons, comparison)
	return nil, remainingArgs
}

func (this *OperationWhere) Run(source *Source) error {
	var rows [][]string
	for _, comparison := range this.Comparisons {
		if err := comparison.PrepareMatch(source); err != nil {
			return err
		}
	}
	rows = append(rows, source.Rows[0])
	for _, row := range source.Rows[1:] {
		for _, comparison := range this.Comparisons {
			if comparison.Match(row) {
				rows = append(rows, row)
				break
			}
		}
	}
	source.Rows = rows
	return nil
}
