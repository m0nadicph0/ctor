package model

type Task struct {
	Name     string   `yaml:"-"`
	Commands []string `yaml:"cmds"`
}
