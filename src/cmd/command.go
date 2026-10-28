package cmd

import (
	"github.com/ajr-cabbage/sssg/src/internal/config"
)

type Command struct {
	Name        string
	Description string
	Help        string
	Callback    func(config.Config, []string) error
}

func GetCommands() map[string]Command {
	commands := make(map[string]Command)
	//TODO: make help command func
	commands["help"] = Command{
		Name:        "help",
		Description: "Display usage information. Try 'sssg help <command>'",
		Callback:    helpCommand,
		Help: `
'sssg help'
List all available commands.
'sssg help <commands...>
Display the detailed help information of one or more commands`,
	}
	commands["init"] = Command{
		Name:        "init",
		Description: "Initialize a project in the current directory.",
		Callback:    initCommand,
		Help: `
'sssg init'
Initialize a project in the current working directory.
If there is a config.json present, it will create required
directories according to that config file.
If there is no config.json, the project will be initialized
with application defaults and a default config.json created.`,
	}
	commands["build"] = Command{
		Name:        "build",
		Description: "Generate a site based on the current configuration.",
		Callback:    buildCommand,
		Help: `
'sssg build'
Convert *.md files from 'source_path' to *.html files in 'target_path' and
copy assets from 'asset_path' to the 'target_path'/assets. Directory
structure of the source_path will be preserved.
This command will verify the configuration matches the current directory
structure. Hint: you can always reinitialize to guaruntee a valid
configuration.`,
	}
	commands["config"] = Command{
		Name:        "config",
		Description: "Modify configuration via CLI.",
		Callback:    configCommand,
		Help: `
'sssg config <subcmd> <args>'
The config command gets and sets current config key-value pairs
and updates config.json. Invalid assignments are noted.
'sssg config get <key>' - Return the value of the provided key.
'sssg config set <key> <value> - Set the key-value pair'`,
	}
	commands["serve"] = Command{
		Name:        "serve",
		Description: "Serve the target directory over http.",
		Callback:    serveCommand,
		Help: `
'sssg serve -<flags> <args>'
Serve the target_path over http.
If no flags are provided, it will be served on the port in config.json.
If -p is supplied, the port will be overridden by the argument.
if -d is supplied, the server will run in detatched mode.`,
	}
	return commands
}
