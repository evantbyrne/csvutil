package lib

import (
	"errors"
	"fmt"
	"strings"
)

type OperationConcat struct{}

func (this *OperationConcat) Construct(source *Source, args []string) (error, []string) {
	return nil, args[1:]
}

func (this *OperationConcat) Run(source *Source) error {
	if source.Previous == nil {
		return errors.New("--concat operation cannot be run on first source.")
	}

	if !ArraysEqual(source.Previous.Rows[0], source.Rows[0]) {
		columns := strings.Join(source.Rows[0], ",")
		columnsPrevious := strings.Join(source.Previous.Rows[0], ",")
		return fmt.Errorf("--concat operation requires both sources to have matching columns. Given sources have columns '%s' and '%s'. Consider using --select operation to resolve.", columnsPrevious, columns)
	}

	source.Rows = append(source.Previous.Rows, source.Rows[1:]...)
	return nil
}
