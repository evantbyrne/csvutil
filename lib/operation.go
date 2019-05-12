package lib

import (
	"fmt"
	"os"
)

type Operation interface {
	Construct(string) error
	Run(*Source) error
}

func MapOperation(name string, arg string) (error, Operation) {
	var operation Operation

	switch name {
	case "except":
		operation = &OperationExcept{}
		break
	case "select":
		operation = &OperationSelect{}
		break
	default:
		fmt.Printf("No operation matching '--%s'.\n", name)
		os.Exit(1)
	}

	err := operation.Construct(arg)
	return err, operation
}
