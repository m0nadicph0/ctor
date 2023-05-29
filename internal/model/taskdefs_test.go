package model

import (
	"testing"
)

func TestTaskDefs_GetDependencies(t *testing.T) {

	taskDefs := &TaskDefs{
		Version: "1",
		Tasks: map[string]*Task{
			"task1": {
				Name: "task1",
				Commands: []*Command{
					{
						Cmd:    "echo task1",
						Task:   "",
						IsTask: false,
					},
				},
				Description:  "sample task",
				Variables:    map[string]any{},
				Dependencies: []string{},
			},
			"task2": {
				Name: "task2",
				Commands: []*Command{
					{
						Cmd:    "echo task2",
						Task:   "",
						IsTask: false,
					},
				},
				Description:  "sample task",
				Variables:    map[string]any{},
				Dependencies: []string{},
			},
			"task3": {
				Name: "task3",
				Commands: []*Command{
					{
						Cmd:    "echo task3",
						Task:   "",
						IsTask: false,
					},
				},
				Description:  "sample task",
				Variables:    map[string]any{},
				Dependencies: []string{"task1", "task2"},
			},
		},
	}

	task, ok := taskDefs.Find("task3")

	if !ok {
		t.Fatalf("expected Find to return a task, but got=%t", ok)
	}

	deps := taskDefs.GetDependencies(task)
	expectedDeps := 2

	if len(deps) != expectedDeps {
		t.Errorf("expected to find %d dependencies, but got %d", expectedDeps, len(deps))
	}
}
