package main

import (
	"fmt"
	"os"
	//"github.com/ajr-cabbage/sssg/src/internal/build"
)

func main() {
	/*conf := build.Config{
		SourcePath:   "./content/docs/",
		TargetPath:   "./target/",
		AssetPath:    "./content/assets/",
		TemplatePath: "./templates/default.html",
	}
	err := conf.Build()
	err := initialize.Initialize()
	if err != nil {
		log.Fatal(err)
		} */

	//helpCommand := flag.NewFlagSet("help", flag.ExitOnError)
	//initCommand := flag.NewFlagSet("init", flag.ExitOnError)
	//buildCommand := flag.NewFlagSet("build", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("sssg - The Small, Shitty, (static) Site Generator")
		fmt.Println("'sssg help' for usage information")
		os.Exit(0)
	}

	switch os.Args[1] {
	case "help":
		fmt.Println("Welcome to sssg - The Small, Shitty, (static) Site Generator")
		fmt.Println("\nCommands: 'sssg <cmd> <flags> <parameters>")
		fmt.Println(" - 'sssg init': Initialize a project in CWD with default configuration.")
		fmt.Println("                If 'config.json' is present, project will use that as default.")
	default:
		fmt.Printf("Unknown command: '%s'\n'sssg help' for usage information\n", os.Args[1])
	}
}
