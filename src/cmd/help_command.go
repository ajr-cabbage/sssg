package cmd

import (
	"fmt"
	"maps"
	"slices"

	"github.com/ajr-cabbage/sssg/src/internal/config"
)

func helpCommand(conf *config.Config, _ []string) error {
	cmds := GetCommands()
	cmdNames := slices.Collect(maps.Keys(cmds))
	slices.Sort(cmdNames)
	fmt.Print("sssg - The Small, Shitty, (Static) Site Generator\n\n")
	fmt.Println("Available Commands:")
	for _, cmdName := range cmdNames {
		fmt.Printf(" '%s' - %s\n\n", cmdName, cmds[cmdName].Description)
	}
	return nil
}
