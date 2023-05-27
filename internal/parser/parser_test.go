package parser

import (
	"bytes"
	"github.com/m0nadicph0/ctor/internal/model"
	"reflect"
	"testing"
)

const TestYamlNoTask = `
version: '1'
`

const TestYamlWithTwoTasks = `
version: '1'

tasks:
  build:
    cmds:
      - go build
  test:
    cmds:
      - go test -v ./...
`

const TestYamlWithThreeTasks = `
version: '1'

tasks:
  default:
    cmds:
      - go build
  test:
    cmds:
      - go test -v ./...
  
  clean:
    cmds:
      - go clean
`

func TestParseTaskDefs(t *testing.T) {

	tests := []struct {
		name    string
		content string
		count   int
		tasks   []*model.Task
		wantErr bool
	}{
		{
			name:    "No Tasks",
			content: TestYamlNoTask,
			count:   0,
			tasks:   []*model.Task{},
			wantErr: false,
		},
		{
			name:    "Two tasks",
			content: TestYamlWithTwoTasks,
			count:   2,
			tasks: []*model.Task{
				{
					Name:     "build",
					Commands: []string{"go build"},
				},
				{
					Name:     "test",
					Commands: []string{"go test -v ./..."},
				},
			},
			wantErr: false,
		},
		{
			name:    "Three tasks",
			content: TestYamlWithThreeTasks,
			count:   3,
			tasks: []*model.Task{
				{
					Name:     "default",
					Commands: []string{"go build"},
				},
				{
					Name:     "test",
					Commands: []string{"go test -v ./..."},
				},
				{
					Name:     "clean",
					Commands: []string{"go clean"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := bytes.NewBuffer([]byte(tt.content))
			taskDefs, err := ParseTaskDefs(data)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseTaskDefs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for _, task := range tt.tasks {
				got := taskDefs.Tasks[task.Name]
				if !reflect.DeepEqual(got, task) {
					t.Errorf("wanted task=%v, but got=%v", task, got)
				}
			}
		})
	}
}

const TestYamlVars = `
version: '1'

vars:
  GREETING: Hello from Taskfile!
  BUILD_NO: 100
  TAG: 0a34ff

tasks:
  default:
    cmds:
      - echo "{{.GREETING}}"
`

func TestParseTaskDefsVars(t *testing.T) {
	data := bytes.NewBuffer([]byte(TestYamlVars))
	taskDefs, err := ParseTaskDefs(data)

	if err != nil {
		t.Fatalf("didn't expect error while parsing yaml, but got %v", err)
	}

	expectedVarsCount := 3
	if len(taskDefs.Variables) != expectedVarsCount {
		t.Errorf("wanted variable count=%d, but got=%d", expectedVarsCount, len(taskDefs.Variables))
	}

	expectedVars := []string{"GREETING", "BUILD_NO", "TAG"}
	expectedVals := map[string]string{
		"GREETING": "Hello from Taskfile!",
		"BUILD_NO": "100",
		"TAG":      "0a34ff",
	}

	for _, expectedVar := range expectedVars {
		variables := taskDefs.GetVars()
		value, ok := variables[expectedVar]

		if !ok {
			t.Errorf("expected to find variable %s, but got=%t", expectedVar, ok)
		}

		expectedValue := expectedVals[expectedVar]

		if expectedValue != value {
			t.Errorf("expected to variable %s to have value=%s, but got=%s", expectedVar, expectedValue, value)
		}
	}

}
