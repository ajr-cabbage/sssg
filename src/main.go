package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ajr-cabbage/sssg/src/cmd"
	"github.com/ajr-cabbage/sssg/src/internal/config"
)

func main() {
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
			err := cmd.Callback(conf, os.Args[2:])
			if err != nil {
				log.Fatal(err)
			}
		} else {
			err := cmd.Callback(conf, []string{})
			if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		fmt.Printf("Unknown command: '%s'\n'sssg help' for usage information\n", os.Args[1])
	}
}
