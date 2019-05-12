package lib

import (
	// "errors"
	"fmt"
	"strings"
)

type Comparison struct {
	Left      string
	LeftType  string
	Operator  string
	Right     string
	RightType string
}

func ConstructComparison(name string, value string) (error, *Comparison) {
	comparison := &Comparison{
		Operator: "=",
	}

	if value == "" {
		return fmt.Errorf("--%s expected comparison.\n", name), nil
	}

	parts := strings.SplitN(value, "=", 2)
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return fmt.Errorf("--%s expected comparison, recieved '%s'.\n", name, value), nil
	}

	if strings.HasPrefix(parts[0], ":") {
		comparison.Left = parts[0][1:]
		comparison.LeftType = "column"
	} else {
		// Forcing left side to be a column simplifies parsing.
		return fmt.Errorf("--%s requires left side of comparison to reference a column (e.g., ':name=value'), recieved '%s'.", name, value), nil
	}

	if strings.HasPrefix(parts[1], ":") {
		comparison.Right = parts[1][1:]
		comparison.RightType = "column"
	} else {
		comparison.Right = parts[1]
		comparison.RightType = "value"
	}

	return nil, comparison
}
