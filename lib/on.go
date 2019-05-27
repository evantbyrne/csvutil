package lib

import (
	"fmt"
	"strings"
)

type On struct {
	ColumnLeft  string
	ColumnRight string
	IndexLeft   int
	IndexRight  int
	Operator    string
}

func (this *On) Match(leftRow []string, rightRow []string) bool {
	switch this.Operator {
	case "==":
		return leftRow[this.IndexLeft] == rightRow[this.IndexRight]
	case "!=":
		return leftRow[this.IndexLeft] != rightRow[this.IndexRight]
	}
	return false
}

func (this *On) PrepareMatch(source *Source) {
	this.IndexRight = source.ColumnIndex(this.ColumnRight)
	this.IndexLeft = source.Previous.ColumnIndex(this.ColumnLeft)
}

func ConstructOn(operation string, columnLeft string, operator string, columnRight string) (error, *On) {
	operator = strings.ToUpper(operator)
	on := &On{
		ColumnLeft:  columnLeft,
		ColumnRight: columnRight,
		Operator:    operator,
	}

	if columnLeft == "" {
		return fmt.Errorf("%s opreration requires a column name for the first argument, which is matched against the previous source.", operation), nil
	}

	if operator != "==" && operator != "!=" {
		return fmt.Errorf("%s opreration requires second argument to be valid operator, '%s' given.", operation, operator), nil
	}

	if columnRight == "" {
		return fmt.Errorf("%s opreration requires a column name for the third argument, which is matched against the current source.", operation), nil
	}

	return nil, on
}
