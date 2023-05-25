package model

import (
	"bytes"
	"fmt"
	"html/template"
)

type Task struct {
	Name        string   `yaml:"-"`
	Commands    []string `yaml:"cmds"`
	Description string   `yaml:"desc"`
}

func (t *Task) GetExpandedCommands(variables map[string]string) ([]string, error) {
	expandedCommands := make([]string, 0)
	for _, command := range t.Commands {
		tmpl, err := template.New("").Parse(command)
		if err != nil {
			return []string{}, fmt.Errorf("failed to parse command [%s]:%v", command, err)
		}
		var buf bytes.Buffer
		_ = tmpl.Execute(&buf, variables)
		expandedCommands = append(expandedCommands, buf.String())
	}
	return expandedCommands, nil
}
