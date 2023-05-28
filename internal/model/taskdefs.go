package model

import (
	"fmt"
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

func (td *TaskDefs) GetVars() map[string]string {
	result := make(map[string]string)
	for key, value := range td.Variables {
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

func (td *TaskDefs) AddVar(key string, value string) {
	if td.Variables == nil {
		td.Variables = make(map[string]any)
	}
	td.Variables[key] = value
}
