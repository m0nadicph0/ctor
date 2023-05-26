package executor

import (
	"github.com/m0nadicph0/ctor/internal/model"
	"github.com/m0nadicph0/ctor/internal/util"
	"io"
	"os"
	"os/exec"
)

type Executor interface {
	Exec(task *model.Task) error
}

type executor struct {
	Stdout   io.Writer
	Stderr   io.Writer
	TaskDefs *model.TaskDefs
}

func NewExecutor(taskDefs *model.TaskDefs) Executor {
	return &executor{
		Stdout:   os.Stdout,
		Stderr:   os.Stderr,
		TaskDefs: taskDefs,
	}
}

func (e *executor) Exec(task *model.Task) error {
	mergedVars := util.MergeVars(e.TaskDefs.Variables, task.Variables)
	expandedCmd, err := task.GetExpandedCommands(mergedVars)

	if err != nil {
		return err
	}

	for _, cmd := range expandedCmd {
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
