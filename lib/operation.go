package lib

import (
	"fmt"
	"os"
)

type Operation interface {
	Construct([]string) (error, []string)
	Run(*Source) error
}

func MapOperation(args []string) (error, Operation, []string) {
	var operation Operation

	switch args[0] {
	case "--except":
		operation = &OperationExcept{}
		break
	case "--select":
		operation = &OperationSelect{}
		break
	case "--where":
		operation = &OperationWhere{}
		break
	default:
		fmt.Printf("No operation matching '%s'.\n", args[0])
		os.Exit(1)
	}

	err, remainingArgs := operation.Construct(args)
	return err, operation, remainingArgs
}
