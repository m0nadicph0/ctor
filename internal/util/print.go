package util

import (
	"fmt"
	"github.com/m0nadicph0/ctor/internal/model"
	"io"
)

func PrintTasks(out io.Writer, tasks []*model.Task) {
	fmt.Fprintln(out, "ctor: Available tasks for this project:")
	for _, task := range tasks {
		fmt.Fprintf(out, "- %s:\t %s\n", task.Name, task.Description)
	}
}
