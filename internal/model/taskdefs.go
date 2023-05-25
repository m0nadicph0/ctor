package model

type TaskDefs struct {
	Version string           `yaml:"version"`
	Tasks   map[string]*Task `yaml:"tasks"`
}

func (td *TaskDefs) Find(name string) (*Task, bool) {
	task, ok := td.Tasks[name]
	return task, ok
}
