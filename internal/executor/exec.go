package executor

import (
	"fmt"
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

func NewExecutorWithWriters(taskDefs *model.TaskDefs, out io.Writer, err io.Writer) Executor {
	return &executor{
		Stdout:   out,
		Stderr:   err,
		TaskDefs: taskDefs,
	}
}

func (e *executor) Exec(task *model.Task) error {
	if task.HasDependency() {
		for _, dependency := range e.TaskDefs.GetDependencies(task) {
			err := e.Exec(dependency)
			if err != nil {
				return fmt.Errorf("failed to execute dependency '%s' for task '%s': %w", dependency.Name, task.Name, err)
			}
		}
	}

	return e.executeTask(task)
}

func (e *executor) executeTask(task *model.Task) error {
	mergedVars := util.MergeVars(e.TaskDefs.GetVars(), task.GetVars())
	expandedCmd, err := task.GetExpandedCommands(mergedVars)

	if err != nil {
		return err
	}

	mergedEnv := util.MergeVars(e.TaskDefs.EnvVars, task.EnvVars)

	for _, cmd := range expandedCmd {
		if cmd.IsTask {
			cmdTask, ok := e.TaskDefs.Find(cmd.Task)
			if ok {
				err := e.executeTask(cmdTask)
				if err != nil {
					return err
				}
			}
		} else {
			err := e.execCmd(cmd.Cmd, util.EnvList(mergedEnv))
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func (e *executor) execCmd(command string, env []string) error {
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = e.Stdout
	cmd.Stderr = e.Stderr
	cmd.Env = env
	return cmd.Run()
}
