package cmd

import (
	"github.com/ajr-cabbage/sssg/src/internal/config"
)

type Command struct {
	Name        string
	Description string
	Callback    func(*config.Config, []string) error
}

func GetCommands() map[string]Command {
	commands := make(map[string]Command)
	//TODO: make help command func
	commands["help"] = Command{
		Name:        "help",
		Description: "Display usage information. Try 'sssg help <command>'",
		Callback:    helpCommand,
	}
	commands["init"] = Command{
		Name:        "init",
		Description: "Initialize a project in the current directory.",
		Callback:    initCommand,
	}
	commands["build"] = Command{
		Name:        "build",
		Description: "Generate a site based on the current configuration.",
		Callback:    buildCommand,
	}
	commands["config"] = Command{
		Name:        "config",
		Description: "Modify configuration via CLI,",
		Callback:    configCommand,
	}
	return commands
}
