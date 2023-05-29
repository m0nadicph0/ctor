package shell

import (
	"fmt"
	"github.com/m0nadicph0/ctor/internal/model"
	"github.com/mattn/go-shellwords"
	"strings"
)

func dryExecTask(args []string, taskDefs *model.TaskDefs) error {
	return nil
}

func execTask(args []string, taskDefs *model.TaskDefs) error {
	return nil
}

func addTask(args []string, taskDefs *model.TaskDefs) error {
	tokens, _ := shellwords.Parse(strings.Join(args, " "))
	fmt.Println(tokens, len(tokens))
	switch len(tokens) {
	case 1:
		taskName := strings.TrimPrefix(tokens[0], "name=")
		taskDefs.AddTask(&model.Task{
			Name: taskName,
		})
	default:
		params := toParamsMap(tokens)
		taskDefs.AddTask(&model.Task{
			Name: params["name"],
			Commands: []*model.Command{
				{
					Cmd:    params["cmd"],
					Task:   "",
					IsTask: false,
				},
			},
			Description: params["desc"],
		})

	}
	return nil
}

func toParamsMap(tokens []string) map[string]string {
	result := make(map[string]string)
	for _, token := range tokens {
		components := strings.Split(token, "=")
		if len(components) == 2 {
			result[components[0]] = components[1]
		}
	}
	return result
}

func listTasks(args []string, taskDefs *model.TaskDefs) error {
	tasks := taskDefs.GetTasks()
	for _, task := range tasks {
		fmt.Println(task.Name)
	}
	return nil
}
