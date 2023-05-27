package util

import (
	"fmt"
	"os"
)

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

func EnvList(env map[string]string) []string {
	envs := os.Environ()
	for k, v := range env {
		envs = append(envs, fmt.Sprintf("%s=%s", k, v))
	}
	return envs
}
