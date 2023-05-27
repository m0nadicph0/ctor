package util

func IsEmpty[T any](a []T) bool {
	return len(a) == 0
}

func SplitArgs(args []string, sep string) ([]string, []string) {
	index := -1
	for i, item := range args {
		if item == sep {
			index = i
			break
		}
	}

	if index == -1 {
		return args, []string{}
	}

	return args[:index], args[index+1:]
}
