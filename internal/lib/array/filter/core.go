package filter

import "strings"

func FilterIsFlag(args []string) []string {
	flags := make([]string, 0, len(args))
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			flags = append(flags, arg)
		}
	}
	return flags
}
