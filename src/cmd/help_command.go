package cmd

import (
	"fmt"

	"github.com/ajr-cabbage/sssg/src/internal/config"
)

func helpCommand(conf *config.Config, _ []string) error {
	fmt.Println("TESTING HELP COMMAND CALLBACK")
	return nil
}
