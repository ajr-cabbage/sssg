package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ajr-cabbage/sssg/src/cmd"
	"github.com/ajr-cabbage/sssg/src/internal/config"
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

	if len(os.Args) < 2 {
		fmt.Println("sssg - The Small, Shitty, (static) Site Generator")
		fmt.Println("'sssg help' for usage information")
		os.Exit(0)
	}
	conf, _, err := config.GetConfig()
	if err != nil {
		fmt.Println(err)
	}
	cmd, ok := cmd.GetCommands()[os.Args[1]]
	if ok {
		if len(os.Args) > 2 {
			err := cmd.Callback(&conf, os.Args[2:])
			if err != nil {
				log.Fatal(err)
			}
		} else {
			err := cmd.Callback(&conf, []string{})
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
