package lib

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Source struct {
	Operations []Operation
	Previous   *Source
	Path       string
	Rows       [][]string
}

func (this *Source) ColumnIndex(key string) (error, int) {
	for i, value := range this.Rows[0] {
		if key == value {
			return nil, i
		}
	}

	return fmt.Errorf("Invalid column '%s' for '%s'.", key, this.Path), -1
}

func (this *Source) MapOperation(args []string) (error, Operation, []string) {
	var operation Operation

	switch args[0] {
	case "--concat":
		operation = &OperationConcat{}
		break
	case "--distinct":
		operation = &OperationDistinct{}
		break
	case "--count":
		operation = &OperationCount{}
		break
	case "--except":
		operation = &OperationExcept{}
		break
	case "--join":
		operation = &OperationJoin{}
		break
	case "--or":
		operation = &OperationOr{}
		break
	case "--select":
		operation = &OperationSelect{}
		break
	case "--sort":
		operation = &OperationSort{}
		break
	case "--values":
		operation = &OperationValues{}
		break
	case "--where":
		operation = &OperationWhere{}
		break
	default:
		return fmt.Errorf("No operation matching '%s'.", args[0]), operation, []string{}
	}

	err, remainingArgs := operation.Construct(this, args)
	return err, operation, remainingArgs
}

func (this *Source) ReadAll() error {
	fh, err := os.Open(this.Path)
	defer fh.Close()
	if err != nil {
		return err
	}

	rows, err := csv.NewReader(fh).ReadAll()
	if err != nil {
		return err
	}

	if len(rows) == 0 {
		return fmt.Errorf("Empty source '%s'.", this.Path)
	}

	this.Rows = rows
	return nil
}

func (this *Source) Run() error {
	if this.Previous != nil {
		this.Previous.Run()
	}

	if err := this.ReadAll(); err != nil {
		return err
	}

	for _, operation := range this.Operations {
		if err := operation.Run(this); err != nil {
			return err
		}
	}

	return nil
}
