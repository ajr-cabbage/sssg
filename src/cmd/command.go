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
		Description: "Display usage information.",
		Callback:    helpCommand,
	}
	//TODO: make cmd/initcallback()
	commands["init"] = Command{
		Name:        "init",
		Description: "Initialize a project in the current directory.",
		Callback:    initCommand,
	}
	//TODO: make cmd/buildcallback()
	commands["build"] = Command{
		Name:        "build",
		Description: "Generate a site in based on the current configuration.",
		Callback:    buildCommand,
	}
	return commands
}
