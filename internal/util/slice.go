package util

func IsEmpty[T any](a []T) bool {
	return len(a) == 0
}
