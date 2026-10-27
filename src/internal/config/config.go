package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	SourcePath   string `json:"source_path"`
	TargetPath   string `json:"target_path"`
	AssetPath    string `json:"asset_path"`
	TemplatePath string `json:"template_path"`
	ThemePath    string `json:"theme_path"`
	Port         int    `json:"port"`
}

func GetConfig() (Config, bool, error) {
	defaultConfig := Config{
		SourcePath:   "./content/docs/",
		TargetPath:   "./target/",
		AssetPath:    "./content/assets/",
		TemplatePath: "./templates/",
		ThemePath:    "./themes/",
		Port:         8080,
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
		return defaultConfig, true, nil
	}
	return defaultConfig, true, err
}

func WriteConfig(c Config) error {
	file, err := os.Create("config.json")
	if err != nil {
		return fmt.Errorf("Error creating config.json: %w\n", err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(c)
	if err != nil {
		return fmt.Errorf("Error writing to config.json: %w\n", err)
	}
	return nil
}

func SetConfig(c Config, key, val string) error {
	switch key {
	case "source_path":
		c.SourcePath = val
		err := WriteConfig(c)
		if err != nil {
			return err
		}
	case "target_path":
		c.TargetPath = val
		err := WriteConfig(c)
		if err != nil {
			return err
		}
	case "asset_path":
		c.AssetPath = val
		err := WriteConfig(c)
		if err != nil {
			return err
		}
	case "template_path":
		c.TemplatePath = val
		err := WriteConfig(c)
		if err != nil {
			return err
		}
	case "theme_path":
		c.ThemePath = val
		err := WriteConfig(c)
		if err != nil {
			return err
		}
	case "port":
		portNum, err := strconv.Atoi(val)
		if err != nil {
			return fmt.Errorf("Invalid port number: %w", err)
		}
		c.Port = portNum
		err = WriteConfig(c)
		if err != nil {
			return err
		}
	default:
		return errors.New("Invalid config field.")
	}
	return nil
}

func GetConfigValue(c Config, key string) any {
	switch key {
	case "source_path":
		return c.SourcePath
	case "target_path":
		return c.TargetPath
	case "asset_path":
		return c.AssetPath
	case "template_path":
		return c.TemplatePath
	case "theme_path":
		return c.ThemePath
	case "port":
		return c.Port
	default:
		return ""
	}
}

func VerifyConfig(c Config) bool {
	_, err := os.Stat("config.json")
	if err != nil {
		return false
	}
	_, err = os.Stat(c.SourcePath)
	if err != nil {
		return false
	}
	_, err = os.Stat(c.TargetPath)
	if err != nil {
		return false
	}
	_, err = os.Stat(c.AssetPath)
	if err != nil {
		return false
	}
	_, err = os.Stat(c.TemplatePath)
	if err != nil {
		return false
	}
	_, err = os.Stat(c.ThemePath)
	if err != nil {
		return false
	}
	return true
}
