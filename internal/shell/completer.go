package shell

import "github.com/chzyer/readline"

var completer = readline.NewPrefixCompleter(
	readline.PcItem("tasks",
		readline.PcItem("ls"),
		readline.PcItem("add",
			readline.PcItem("name="),
			readline.PcItem("desc="),
			readline.PcItem("cmd="),
		),
		readline.PcItem("exec"),
		readline.PcItem("dry"),
	),
	readline.PcItem("vars"),
	readline.PcItem("exit"),
	readline.PcItem("help"),
	readline.PcItem("save"),
)
