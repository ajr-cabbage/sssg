package cmd

import (
	"errors"

	"github.com/ajr-cabbage/sssg/src/internal/config"
	"github.com/ajr-cabbage/sssg/src/internal/theme"
)

func themeCommand(conf config.Config, args []string) error {
	if len(args) == 0 {
		return errors.New("Insufficient arguments. See 'sssg help theme' for usage.")
	}
	if len(args) > 1 {
		return errors.New("Too many arguments. See 'sssg help theme' for usage.")
	}
	err := theme.UpdateTheme(conf, args[0])
	if err != nil {
		return err
	}
	return nil
}
