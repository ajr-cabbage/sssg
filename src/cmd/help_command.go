package cmd

import (
	"fmt"
	"maps"
	"slices"

	"github.com/ajr-cabbage/sssg/src/internal/config"
)

func helpCommand(conf config.Config, args []string) error {
	cmds := GetCommands()
	if len(args) == 0 {
		cmdNames := slices.Collect(maps.Keys(cmds))
		slices.Sort(cmdNames)
		fmt.Print("sssg - The Small, Shitty, (Static) Site Generator\n\n")
		fmt.Println("Available Commands:")
		for _, cmdName := range cmdNames {
			fmt.Printf(" '%s' - %s\n\n", cmdName, cmds[cmdName].Description)
		}
		return nil
	}
	for _, arg := range args {
		cmd, ok := cmds[arg]
		if ok {
			fmt.Println(cmd.Help)
		} else {
			fmt.Printf("Unknown command: '%s'. 'sssg help' to list available commands.\n", arg)
		}
	}
	return nil
}
