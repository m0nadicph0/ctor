package model

type TaskDefs struct {
	Version   string            `yaml:"version"`
	Variables map[string]string `yaml:"vars"`
	Tasks     map[string]*Task  `yaml:"tasks"`
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

func (td TaskDefs) GetDependencies(task *Task) []*Task {
	dependencies := make([]*Task, 0)
	for _, dep := range task.Dependencies {
		dt, ok := td.Find(dep)
		if ok {
			dependencies = append(dependencies, dt)
		}
	}
	return dependencies
}
