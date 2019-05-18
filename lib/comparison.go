package lib

import (
	"fmt"
)

type Comparison struct {
	Column   string
	Operator string
	Values   []string
}

func (this *Comparison) Match(source *Source, row []string) bool {
	columnIndex := source.ColumnIndex(this.Column)
	switch this.Operator {
	case "==":
		return row[columnIndex] == this.Values[0]
	case "!=":
		return row[columnIndex] != this.Values[0]
	}
	return false
}

func ConstructComparison(operation string, column string, operator string, args []string) (error, *Comparison, []string) {
	var remainingArgs []string

	comparison := &Comparison{
		Column:   column,
		Operator: operator,
		Values:   []string{},
	}

	if column == "" {
		return fmt.Errorf("%s opreration requires a column name for the first argument.", operation), nil, []string{}
	}

	if len(args) < 1 {
		return fmt.Errorf("%s opreration requires at least three arguments", operation), nil, []string{}
	}

	if operator == "==" || operator == "!=" {
		remainingArgs = args[1:]
		comparison.Values = []string{args[0]}
	} else {
		return fmt.Errorf("%s opreration requires second argument to be valid operator, '%s' given.", operation, operator), nil, []string{}
	}

	return nil, comparison, remainingArgs
}
