package model

import (
	"fmt"
	"os/exec"
)

type TaskDefs struct {
	Version   string            `yaml:"version"`
	Variables map[string]any    `yaml:"vars"`
	EnvVars   map[string]string `yaml:"env"`
	Tasks     map[string]*Task  `yaml:"tasks"`
}

func (td *TaskDefs) Find(name string) (*Task, bool) {
	task, ok := td.Tasks[name]
	return task, ok
}

func (td *TaskDefs) FindByAlias(alias string) (*Task, bool) {
	for _, task := range td.GetTasks() {
		if task.HasAlias(alias) {
			return task, true
		}
	}
	return nil, false
}

func (td *TaskDefs) GetTasks() []*Task {
	tasks := make([]*Task, 0)
	for _, task := range td.Tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

func (td *TaskDefs) GetTasksWithDesc() []*Task {
	tasks := make([]*Task, 0)
	for _, task := range td.Tasks {
		if len(task.Description) > 0 {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

func (td *TaskDefs) GetDependencies(task *Task) []*Task {
	dependencies := make([]*Task, 0)
	for _, dep := range task.Dependencies {
		dt, ok := td.Find(dep)
		if ok {
			dependencies = append(dependencies, dt)
		}
	}
	return dependencies
}

func (td TaskDefs) GetVars() map[string]string {
	result := make(map[string]string)
	for key, value := range td.Variables {
		switch value.(type) {
		case string:
			strKey := fmt.Sprintf("%v", key)
			result[strKey] = fmt.Sprintf("%v", value)
		case map[any]any:
			strKey := fmt.Sprintf("%v", key)
			result[strKey] = shellExpand(value.(map[any]any))
		default:
			strKey := fmt.Sprintf("%v", key)
			result[strKey] = fmt.Sprintf("%d", value)
		}
	}
	return result
}

func shellExpand(value map[any]any) string {
	shVal := toStrMap(value)
	cmd := shVal["sh"]

	return shellExec(cmd)
}

func toStrMap(dynamic map[any]any) map[string]string {
	result := make(map[string]string)
	for key, value := range dynamic {
		sKey := key.(string)
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

	return string(output)
}
