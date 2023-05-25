package util

import (
	"fmt"
	"github.com/m0nadicph0/ctor/internal/model"
	"io"
)

//func PrintTaskNames(out io.Writer, tasks []*model.Task) {
//	fmt.Fprintln(out, "ctor: Available tasks for this project:")
//	for _, task := range tasks {
//		fmt.Fprintf(out, "- %s\n", task.Name)
//	}
//}

func PrintTasks(out io.Writer, tasks []*model.Task) {
	fmt.Fprintln(out, "ctor: Available tasks for this project:")
	for _, task := range tasks {
		fmt.Fprintf(out, "- %s:\t %s\n", task.Name, task.Description)
	}
}
