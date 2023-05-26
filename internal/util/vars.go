package util

func MergeVars(v1 map[string]string, v2 map[string]string) map[string]string {
	merged := make(map[string]string)
	for k, v := range v1 {
		merged[k] = v
	}
	for k, v := range v2 {
		merged[k] = v
	}
	return merged
}
