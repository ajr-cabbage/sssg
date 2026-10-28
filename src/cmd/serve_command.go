package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/ajr-cabbage/sssg/src/internal/config"
	"github.com/ajr-cabbage/sssg/src/internal/serve"
)

func serveCommand(conf config.Config, _ []string) error {
	_, err := os.Stat(conf.TargetPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return errors.New("Target path does not exist.")
		} else {
			return fmt.Errorf("Error when checking target path: %s\n", err)
		}
	}
	err = serve.Serve(conf.TargetPath, conf.Port)
	if err != nil {
		return err
	}
	return nil
}
