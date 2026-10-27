package initialize

import (
	"fmt"
	"os"

	"github.com/ajr-cabbage/sssg/src/internal/config"
)

func Initialize() error {
	conf, isDefault, err := config.GetConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	if isDefault {
		fmt.Println("Initializing with default configuration...")
		err = config.WriteConfig(conf)
		if err != nil {
			return err
		}
	} else {
		fmt.Println("Initializing with configuration from 'config.json'...")
	}
	err = os.MkdirAll(conf.SourcePath, 0744)
	if err != nil {
		return err
	}
	err = os.MkdirAll(conf.TargetPath, 0744)
	if err != nil {
		return err
	}
	err = os.MkdirAll(conf.AssetPath, 0744)
	if err != nil {
		return err
	}
	err = os.MkdirAll(conf.TemplatePath, 0744)
	if err != nil {
		return err
	}
	err = os.MkdirAll(conf.ThemePath, 0744)
	if err != nil {
		return err
	}
	fmt.Println("Complete!")
	return nil
}
