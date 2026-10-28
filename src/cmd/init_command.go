package cmd

import (
	"github.com/ajr-cabbage/sssg/src/internal/config"
	"github.com/ajr-cabbage/sssg/src/internal/initialize"
)

func initCommand(_ config.Config, _ []string) error {
	err := initialize.Initialize()
	return err
}
