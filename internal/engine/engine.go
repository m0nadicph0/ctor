package engine

import (
	"github.com/m0nadicph0/ctor/internal/executor"
	"github.com/m0nadicph0/ctor/internal/model"
	"github.com/m0nadicph0/ctor/internal/util"
)

type Engine struct {
	executor executor.Executor
	taskDefs *model.TaskDefs
}

func NewEngine(exec executor.Executor, td *model.TaskDefs) *Engine {
	return &Engine{executor: exec, taskDefs: td}
}

func (e Engine) Start(args []string) error {
	if len(args) == 0 {
		args = append(args, "default")
	}

	for _, arg := range args {
		task, ok := e.taskDefs.Find(arg)
		if !ok {
			util.ErrExitF(1, "task \"%s\" not found\n", arg)
		}

		expandedCmd, err := task.GetExpandedCommands(e.taskDefs.Variables)

		if err != nil {
			return err
		}

		err = e.executor.Exec(expandedCmd)

		if err != nil {
			return err
		}
	}
	return nil
}
