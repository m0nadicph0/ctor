package executor

import (
	"bytes"
	"github.com/m0nadicph0/ctor/internal/model"
	"testing"
)

func Test_executor_Exec(t *testing.T) {
	taskDefs := &model.TaskDefs{
		Version: "1",
		Tasks: map[string]*model.Task{
			"task1": &model.Task{
				Name:         "task1",
				Commands:     []string{"echo task1"},
				Description:  "sample task",
				Variables:    make(map[string]any),
				Dependencies: []string{},
			},
			"task2": &model.Task{
				Name:         "task2",
				Commands:     []string{"echo task2"},
				Description:  "sample task",
				Variables:    make(map[string]any),
				Dependencies: []string{"task1"},
			},
			"task3": &model.Task{
				Name:         "task3",
				Commands:     []string{"echo task3"},
				Description:  "sample task",
				Variables:    make(map[string]any),
				Dependencies: []string{"task2"},
			},
		},
	}
	out := new(bytes.Buffer)
	errOut := new(bytes.Buffer)
	executor := NewExecutorWithWriters(taskDefs, out, errOut)
	task, _ := taskDefs.Find("task3")
	err := executor.Exec(task)
	if err != nil {
		t.Fatalf("didn't expect errors while executing dependent tasks, but got=%v", err)
	}

	expectedOutput := `task1
task2
task3
`

	if out.String() != expectedOutput {
		t.Errorf("wandted output = %s, but got %s", expectedOutput, out.String())
	}
}
