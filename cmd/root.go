package cmd

import (
	"fmt"
	engine "github.com/m0nadicph0/ctor/internal/engine"
	"github.com/m0nadicph0/ctor/internal/executor"
	"github.com/m0nadicph0/ctor/internal/model"
	"github.com/m0nadicph0/ctor/internal/parser"
	"github.com/m0nadicph0/ctor/internal/util"
	"github.com/m0nadicph0/ctor/version"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "ctor",
	Short: "Ctor is a tool for executing tasks defined in a Ctorfile.yml.",
	Long: `
Ctor is a command-line tool that reads task definitions from a Ctorfile.yml 
and executes the specified tasks. It provides a convenient way to automate 
build processes, manage dependencies between tasks, and perform common 
build requirements.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		showVersion, _ := cmd.Flags().GetBool("version")

		if showVersion {
			fmt.Fprintln(os.Stdout, "Ctor version:", version.GetVersion())
			os.Exit(0)
		}

		ctorFile, _ := cmd.Flags().GetString("ctorfile")
		argsSep, _ := cmd.Flags().GetString("args-sep")

		cf, err := os.Open(ctorFile)
		if err != nil {
			return err
		}
		taskDefs, err := parser.ParseTaskDefs(cf)

		if err != nil {
			return err
		}

		listAll, _ := cmd.Flags().GetBool("list-all")

		if listAll {
			model.PrintTasks(os.Stdout, taskDefs.GetTasks())
			os.Exit(0)
		}

		list, _ := cmd.Flags().GetBool("list")

		if list {
			withDesc := taskDefs.GetTasksWithDesc()
			if util.IsEmpty(withDesc) {
				util.WarnExit("ctor: No tasks with description available. Try --list-all to list all tasks\n")
			}
			model.PrintTasks(os.Stdout, withDesc)
			os.Exit(0)
		}

		eng := engine.NewEngine(executor.NewExecutor(taskDefs), taskDefs)
		core, cliArgs := util.SplitArgs(args, argsSep)
		taskDefs.AddVar("CLI_ARGS", strings.Join(cliArgs, " "))

		return eng.Start(core)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("ctorfile", "c", "Ctorfile.yaml", "Choose which Ctorfile to run")
	rootCmd.Flags().BoolP("verbose", "v", false, "Enables verbose mode")
	rootCmd.Flags().BoolP("help", "h", false, "Shows usage message")
	rootCmd.Flags().BoolP("list", "l", false, "Lists tasks with description of current Ctorfile")
	rootCmd.Flags().BoolP("list-all", "a", false, "Lists tasks with or without a description")
	rootCmd.Flags().StringP("args-sep", "S", "__", "CLI args separator")
	rootCmd.Flags().Bool("version", false, "Show version")

}
