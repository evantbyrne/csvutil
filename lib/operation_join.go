package lib

import (
	"errors"
)

type OperationJoin struct {
	Comparisons []*On
}

func (this *OperationJoin) Construct(source *Source, args []string) (error, []string) {
	if len(args) < 4 {
		return errors.New("--join operation requires three arguments."), []string{}
	}
	err, on := ConstructOn("--join", args[1], args[2], args[3])
	if err != nil {
		return err, []string{}
	}
	this.Comparisons = append(this.Comparisons, on)
	return nil, args[4:]
}

func (this *OperationJoin) Run(source *Source) error {
	var rows [][]string
	header := append(source.Previous.Rows[0], source.Rows[0]...)
	rows = append(rows, header)
	for _, on := range this.Comparisons {
		on.PrepareMatch(source)
	}
	for _, rowRight := range source.Rows[1:] {
		for _, rowLeft := range source.Previous.Rows[1:] {
			for _, on := range this.Comparisons {
				if on.Match(rowLeft, rowRight) {
					row := append(rowLeft, rowRight...)
					rows = append(rows, row)
					break
				}
			}
		}
	}
	source.Rows = rows
	return nil
}
