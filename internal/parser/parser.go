package parser

import (
	"github.com/m0nadicph0/ctor/internal/model"
	"gopkg.in/yaml.v3"
	"io"
)

func ParseTaskDefs(r io.Reader) (*model.TaskDefs, error) {
	data, err := io.ReadAll(r)

	if err != nil {
		return nil, err
	}
	var taskDefs model.TaskDefs

	err = yaml.Unmarshal(data, &taskDefs)

	if err != nil {
		return nil, err
	}

	for taskName, task := range taskDefs.Tasks {
		task.Name = taskName
	}

	return &taskDefs, nil
}
