package shell

import (
	"errors"
	"fmt"
	"github.com/m0nadicph0/ctor/internal/model"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type ShCmdFn func(args []string, taskDefs *model.TaskDefs) error

var taskCmdMap map[string]ShCmdFn = map[string]ShCmdFn{
	"ls":   listTasks,
	"add":  addTask,
	"exec": execTask,
	"dry":  dryExecTask,
}

var cmdMap = map[string]ShCmdFn{
	"exit": func(args []string, taskDefs *model.TaskDefs) error {
		os.Exit(1)
		return nil
	},
	"help": func(args []string, taskDefs *model.TaskDefs) error {
		fmt.Println("Help here")
		return nil
	},
	"tasks": func(args []string, taskDefs *model.TaskDefs) error {
		return LookupAndExec(args, taskCmdMap, taskDefs)
	},
	"save": saveTaskDef,
}

func saveTaskDef(args []string, defs *model.TaskDefs) error {

	if len(args) == 0 {
		return errors.New("missing required argument: file name")
	}

	data, err := yaml.Marshal(defs)
	if err != nil {
		return err
	}

	indentedYAML := strings.ReplaceAll(string(data), "    ", "  ")

	err = os.WriteFile(args[0], []byte(indentedYAML), 0644)
	if err != nil {
		return err
	}
	return nil
}

func LookupAndExec(args []string, funcMap map[string]ShCmdFn, taskDefs *model.TaskDefs) error {
	cmd := args[0]
	fn, ok := funcMap[cmd]
	if !ok {
		return fmt.Errorf("command %s not found", cmd)
	}

	return fn(args[1:], taskDefs)
}
