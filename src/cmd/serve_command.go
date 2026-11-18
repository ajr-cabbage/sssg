package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/ajr-cabbage/sssg/src/internal/config"
	"github.com/ajr-cabbage/sssg/src/internal/serve"
)

func serveCommand(conf config.Config, args []string) error {
	_, err := os.Stat(conf.TargetPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return errors.New("Target path does not exist.")
		} else {
			return fmt.Errorf("Error when checking target path: %w\n", err)
		}
	}
	var port int
	switch len(args) {
	case 0:
		port = conf.Port
	case 2:
		if args[0] == "-p" {
			portInt, err := strconv.Atoi(args[1])
			if err != nil {
				return fmt.Errorf("Invalid port '%s': %w", args[1], err)
			}
			port = portInt
		} else {
			return errors.New("Invalid parameters. 'sssg help serve' for usage.")
		}
	default:
		return errors.New("Invalid parameters. 'sssg help serve' for usage.")
	}
	err = serve.Serve(conf.TargetPath, port)
	if err != nil {
		return err
	}
	return nil
}
