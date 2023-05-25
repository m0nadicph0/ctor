package executor

import (
	"io"
	"os"
	"os/exec"
)

type Executor interface {
	Exec(commands []string) error
}

type executor struct {
	Stdout io.Writer
	Stderr io.Writer
}

func NewExecutor() Executor {
	return &executor{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
}

func (e *executor) Exec(commands []string) error {
	for _, cmd := range commands {
		err := e.execCmd(cmd)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *executor) execCmd(command string) error {
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
