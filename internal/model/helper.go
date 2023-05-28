package model

import (
	"os/exec"
	"strings"
)

func shellExpand(value map[string]any) string {
	shVal := toStrMap(value)
	cmd := shVal["sh"]

	return shellExec(cmd)
}

func toStrMap(dynamic map[string]any) map[string]string {
	result := make(map[string]string)
	for key, value := range dynamic {
		sKey := key
		sValue := value.(string)
		result[sKey] = sValue
	}
	return result
}

func shellExec(cmdStr string) string {
	cmd := exec.Command("sh", "-c", cmdStr)

	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	return strings.ReplaceAll(string(output), "\n", "")
}
