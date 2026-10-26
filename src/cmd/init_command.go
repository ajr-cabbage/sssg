package cmd

import (
	"fmt"

	"github.com/ajr-cabbage/sssg/src/internal/config"
)

func initCommand(conf *config.Config, _ []string) error {
	fmt.Println("TESTING HELP INIT CALLBACK")
	return nil
}
