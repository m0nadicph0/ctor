package model

type Task struct {
	Name        string   `yaml:"-"`
	Commands    []string `yaml:"cmds"`
	Description string   `yaml:"desc"`
}
