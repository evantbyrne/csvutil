package lib

import (
	"fmt"
	"os"
	"strings"
)

func ArgList(args []string) *Source {
	var source *Source

	for _, arg := range args {

		// Operations
		if source != nil && strings.HasPrefix(arg, "--") {
			parts := strings.SplitN(arg[2:], "=", 2)
			value := ""
			if len(parts) == 2 {
				value = parts[1]
			}
			err, operation := MapOperation(parts[0], value)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			source.Operations = append(source.Operations, operation)
		} else {
			// Sources
			if source == nil {
				source = &Source{
					Path: arg,
				}
			} else {
				source = &Source{
					Previous: source,
					Path:     arg,
				}
			}
		}
	}

	return source
}
