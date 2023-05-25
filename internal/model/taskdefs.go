package model

type TaskDefs struct {
	Version string           `yaml:"version"`
	Tasks   map[string]*Task `yaml:"tasks"`
}

func (td *TaskDefs) Find(name string) (*Task, bool) {
	task, ok := td.Tasks[name]
	return task, ok
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
