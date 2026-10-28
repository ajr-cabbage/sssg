package cmd

import (
	"errors"
	"fmt"

	"github.com/ajr-cabbage/sssg/src/internal/config"
)

func configCommand(conf config.Config, args []string) error {
	if len(args) < 2 {
		return errors.New("Insufficient arguments. See 'sssg help config' for usage.")
	}
	switch args[0] {
	case "get":
		val := config.GetConfigValue(conf, args[1])
		if val == "" {
			return fmt.Errorf("Invalid config key: '%s'", args[1])
		}
		fmt.Println(args[1], ": ", val)
	case "set":
		if len(args) != 3 {
			return errors.New("Wrong number of arguments. See 'sssg help config' for usage.")
		}
		err := config.SetConfig(conf, args[1], args[2])
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("Invalid command: '%s'. See 'sssg help config' for usage.", args[0])
	}

	return nil
}
