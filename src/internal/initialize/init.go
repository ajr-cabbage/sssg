package initialize

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Config struct {
	SourcePath   string `json:"source_path"`
	TargetPath   string `json:"target_path"`
	AssetPath    string `json:"asset_path"`
	TemplatePath string `json:"template_path"`
	ThemePath    string `json:"theme_path"`
}

func Initialize() error {
	conf, isDefault, err := getConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	if isDefault {
		fmt.Println("Initializing with default configuration...")
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

func getConfig() (Config, bool, error) {
	defaultConfig := Config{
		SourcePath:   "./content/docs/",
		TargetPath:   "./target/",
		AssetPath:    "./content/assets/",
		TemplatePath: "./templates/",
		ThemePath:    "./themes/",
	}

	_, err := os.Stat("config.json")
	if err == nil {
		jsonBytes, err := os.ReadFile("config.json")
		if err != nil {
			return defaultConfig, true, fmt.Errorf("Error reading config.json: %w\n", err)
		}
		var conf Config
		err = json.Unmarshal(jsonBytes, &conf)
		if err != nil {
			return defaultConfig, true, fmt.Errorf("Error unmarshalling config.json: %w\n", err)
		}
		return conf, false, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		file, err := os.Create("config.json")
		if err != nil {
			return defaultConfig, true, fmt.Errorf("Error creating config.json: %w\n", err)
		}
		defer file.Close()
		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		err = encoder.Encode(defaultConfig)
		if err != nil {
			return defaultConfig, true, fmt.Errorf("Error writing to config.json: %w\n", err)
		}
		return defaultConfig, true, nil
	}
	return defaultConfig, true, err
}
