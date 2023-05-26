package model

import (
	"bytes"
	"fmt"
	"github.com/m0nadicph0/ctor/internal/builtins"
	"html/template"
	"io"
)

type Task struct {
	Name        string            `yaml:"-"`
	Commands    []string          `yaml:"cmds"`
	Description string            `yaml:"desc"`
	Variables   map[string]string `yaml:"vars"`
}

func (t *Task) GetExpandedCommands(variables map[string]string) ([]string, error) {
	expandedCommands := make([]string, 0)
	for _, command := range t.Commands {
		tmpl, err := template.New("").Funcs(template.FuncMap(builtins.BuiltinFunctions)).Parse(command)
		if err != nil {
			return []string{}, fmt.Errorf("failed to parse command [%s]:%v", command, err)
		}
		var buf bytes.Buffer
		_ = tmpl.Execute(&buf, variables)
		expandedCommands = append(expandedCommands, buf.String())
	}
	return expandedCommands, nil
}

func (t *Task) String() string {
	return fmt.Sprintf("- %s:\t %s", t.Name, t.Description)
}

func PrintTasks(out io.Writer, tasks []*Task) {
	fmt.Fprintln(out, "ctor: Available tasks for this project:")
	for _, task := range tasks {
		fmt.Fprintf(out, "- %s:\t %s\n", task.Name, task.Description)
	}
}
