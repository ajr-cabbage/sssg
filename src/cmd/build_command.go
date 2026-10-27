package cmd

import (
	"github.com/ajr-cabbage/sssg/src/internal/build"
	"github.com/ajr-cabbage/sssg/src/internal/config"
)

func buildCommand(conf *config.Config, _ []string) error {
	err := build.Build(conf)
	return err
}
