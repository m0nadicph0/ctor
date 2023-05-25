package cmd

import (
	engine "github.com/m0nadicph0/ctor/internal/engine"
	"github.com/m0nadicph0/ctor/internal/executor"
	"github.com/m0nadicph0/ctor/internal/parser"
	"os"

	"github.com/spf13/cobra"
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
		ctorFile, _ := cmd.Flags().GetString("ctorfile")

		cf, err := os.Open(ctorFile)
		if err != nil {
			return err
		}
		taskDefs, err := parser.ParseTaskDefs(cf)

		if err != nil {
			return err
		}

		eng := engine.NewEngine(executor.NewExecutor(), taskDefs)

		return eng.Start(args)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("ctorfile", "c", "Ctorfile.yaml", "choose which Ctorfile to run")
	rootCmd.Flags().BoolP("verbose", "v", false, "enables verbose mode")
	rootCmd.Flags().BoolP("help", "h", false, "shows usage message")

}
