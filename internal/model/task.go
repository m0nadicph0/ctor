package model

import (
	"bytes"
	"fmt"
	"github.com/logrusorgru/aurora/v4"
	"github.com/m0nadicph0/ctor/internal/builtins"
	"github.com/olekukonko/tablewriter"
	"io"
	"strings"
	"text/template"
)

type Task struct {
	Name         string            `yaml:"-"`
	Commands     []*Command        `yaml:"cmds"`
	Description  string            `yaml:"desc"`
	Variables    map[string]any    `yaml:"vars"`
	EnvVars      map[string]string `yaml:"env"`
	Dependencies []string          `yaml:"deps"`
	Aliases      []string          `yaml:"aliases"`
}

func (t *Task) GetExpandedCommands(variables map[string]string) ([]*Command, error) {
	expandedCommands := make([]*Command, 0)
	for _, command := range t.Commands {
		if command.IsTask {
			expandedCommands = append(expandedCommands, &Command{
				Cmd:    "",
				Task:   command.Task,
				IsTask: true,
			})
		} else {
			tmpl, err := template.New("").Funcs(template.FuncMap(builtins.BuiltinFunctions)).Parse(command.Cmd)
			if err != nil {
				return nil, fmt.Errorf("failed to parse command [%s]:%v", command, err)
			}
			var buf bytes.Buffer
			_ = tmpl.Execute(&buf, variables)
			expandedCommands = append(expandedCommands, &Command{
				Cmd:    buf.String(),
				Task:   "",
				IsTask: false,
			})
		}

	}
	return expandedCommands, nil
}

func (t *Task) String() string {
	return fmt.Sprintf("- %s:\t %s", t.Name, t.Description)
}

func (t *Task) HasDependency() bool {
	return len(t.Dependencies) != 0
}

func (t *Task) HasAlias(alias string) bool {
	for _, a := range t.Aliases {
		if a == alias {
			return true
		}
	}
	return false
}

func (t *Task) GetVars() map[string]string {
	result := make(map[string]string)
	for key, value := range t.Variables {
		switch value.(type) {
		case string:
			strKey := fmt.Sprintf("%v", key)
			result[strKey] = fmt.Sprintf("%v", value)
		case map[string]any:
			strKey := fmt.Sprintf("%v", key)
			result[strKey] = shellExpand(value.(map[string]any))
		case float64:
			strKey := fmt.Sprintf("%v", key)
			result[strKey] = fmt.Sprintf("%0.1f", value)
		default:
			strKey := fmt.Sprintf("%v", key)
			result[strKey] = fmt.Sprintf("%v", value)
		}
	}
	return result
}

func PrintTasks(out io.Writer, tasks []*Task) {
	fmt.Fprintln(out, "ctor: Available tasks for this project:")
	table := tablewriter.NewWriter(out)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetAutoWrapText(false)
	for _, task := range tasks {
		table.Append([]string{aurora.Yellow("-").String(), aurora.Green(task.Name).String(), task.Description, aurora.Cyan(getAliasStr(task)).String()})
	}
	table.Render()
}

func getAliasStr(task *Task) string {
	if len(task.Aliases) == 0 {
		return ""
	}
	return fmt.Sprintf("[aliases: %s]", strings.Join(task.Aliases, ", "))
}
