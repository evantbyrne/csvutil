package lib

import (
	"strings"
)

func ArgList(args []string, source *Source) (error, *Source) {
	if source != nil && strings.HasPrefix(args[0], "--") {
		// Operations
		err, operation, remainingArgs := source.MapOperation(args)
		if err != nil {
			return err, nil
		}
		source.Operations = append(source.Operations, operation)
		args = remainingArgs
	} else {
		// Sources
		if source == nil {
			source = &Source{
				Path: args[0],
			}
		} else {
			source = &Source{
				Previous: source,
				Path:     args[0],
			}
		}
		args = args[1:]
	}

	if len(args) > 0 {
		return ArgList(args, source)
	}

	return nil, source
}
