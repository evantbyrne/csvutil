package lib

import (
	"fmt"
	"strings"
)

type Comparison struct {
	Column      string
	ColumnIndex int
	Operator    string
	Values      []string
}

func (this *Comparison) Match(source *Source, row []string) bool {
	switch this.Operator {
	case "==":
		return row[this.ColumnIndex] == this.Values[0]
	case "!=":
		return row[this.ColumnIndex] != this.Values[0]
	case "IN":
		for _, value := range strings.Split(this.Values[0], ",") {
			if row[this.ColumnIndex] == value {
				return true
			}
		}
	case "NOT_IN":
		for _, value := range strings.Split(this.Values[0], ",") {
			if row[this.ColumnIndex] == value {
				return false
			}
		}
		return true
	}
	return false
}

func (this *Comparison) PrepareMatch(source *Source) error {
	err, columnIndex := source.ColumnIndex(this.Column)
	if err != nil {
		return err
	}
	this.ColumnIndex = columnIndex
	return nil
}

func ConstructComparison(operation string, column string, operator string, args []string) (error, *Comparison, []string) {
	var remainingArgs []string
	operator = strings.ToUpper(operator)
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

	if operator == "==" || operator == "!=" || operator == "IN" || operator == "NOT_IN" {
		remainingArgs = args[1:]
		comparison.Values = []string{args[0]}
	} else {
		return fmt.Errorf("%s opreration requires second argument to be valid operator, '%s' given.", operation, operator), nil, []string{}
	}

	return nil, comparison, remainingArgs
}
