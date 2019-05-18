package lib

import (
	"errors"
)

type OperationOr struct{}

func (this *OperationOr) Construct(source *Source, args []string) (error, []string) {
	var lastWhere *OperationWhere

	for i := len(source.Operations) - 1; i >= 0; i-- {
		switch source.Operations[i].(type) {
		case *OperationOr:
			break
		case *OperationWhere:
			lastWhere = source.Operations[i].(*OperationWhere)
			break
		default:
			return errors.New("--or operation must follow either a '--where' or another '--or' operation."), []string{}
		}
		if lastWhere != nil {
			break
		}
	}

	if lastWhere == nil {
		return errors.New("--or operation must follow either a '--where' or another '--or' operation."), []string{}
	}

	if len(args) < 4 {
		return errors.New("--or operation requires three arguments."), []string{}
	}
	err, comparison, remainingArgs := ConstructComparison("--or", args[1], args[2], args[3:])
	if err != nil {
		return err, []string{}
	}
	lastWhere.Comparisons = append(lastWhere.Comparisons, comparison)
	source.Operations[len(source.Operations)-1] = lastWhere
	return nil, remainingArgs
}

func (this *OperationOr) Run(source *Source) error {
	return nil
}
