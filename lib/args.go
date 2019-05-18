package lib

import (
	"fmt"
	"os"
	"strings"
)

func ArgList(args []string, source *Source) *Source {
	if source != nil && strings.HasPrefix(args[0], "--") {
		// Operations
		err, operation, remainingArgs := MapOperation(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
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

	return source
}
