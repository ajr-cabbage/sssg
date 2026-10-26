package cmd

import (
	"fmt"

	"github.com/ajr-cabbage/sssg/src/internal/config"
)

func buildCommand(conf *config.Config, _ []string) error {
	fmt.Println("TESTING BUILD COMMAND CALLBACK")
	return nil
}
