package shell

import (
	"fmt"
	"github.com/chzyer/readline"
	"github.com/m0nadicph0/ctor/internal/model"
	"github.com/m0nadicph0/ctor/internal/parser"
	"io"
	"log"
	"os"
	"strings"
)

const ArgsSep = " "
const Prompt = "\033[31m»\033[0m "

func filterInput(r rune) (rune, bool) {
	switch r {
	// block CtrlZ feature
	case readline.CharCtrlZ:
		return r, false
	}
	return r, true
}

type Shell struct {
	taskDefs *model.TaskDefs
	ctorFile string
}

func NewShell(fileName string) *Shell {
	return &Shell{
		taskDefs: new(model.TaskDefs),
		ctorFile: fileName,
	}
}

func (sh *Shell) process(line string) error {
	args := strings.Split(line, ArgsSep)
	return LookupAndExec(args, cmdMap, sh.taskDefs)
}

func (sh *Shell) Start() error {

	ctrFile, err := os.Open(sh.ctorFile)
	if err != nil {
		fmt.Printf("ERROR: failed to open %s: %v", sh.ctorFile, err)
	} else {
		sh.taskDefs, err = parser.ParseTaskDefs(ctrFile)
		if err != nil {
			fmt.Printf("ERROR: parsing %s: %v", sh.ctorFile, err)
			sh.taskDefs = new(model.TaskDefs)
		}
		fmt.Println("Loaded", sh.ctorFile)

	}

	l, err := readline.NewEx(&readline.Config{
		Prompt:          Prompt,
		HistoryFile:     "/tmp/readline.tmp",
		AutoComplete:    completer,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",

		HistorySearchFold:   true,
		FuncFilterInputRune: filterInput,
	})
	if err != nil {
		return err
	}
	defer l.Close()
	l.CaptureExitSignal()

	log.SetOutput(l.Stderr())
	for {
		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		err = sh.process(line)
		if err != nil {
			fmt.Fprintln(os.Stderr, "ERROR:", err)
		}

	}

	return nil
}
