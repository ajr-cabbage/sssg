package theme

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/ajr-cabbage/sssg/src/internal/config"
	"github.com/ajr-cabbage/sssg/src/internal/copy"
)

func UpdateTheme(c config.Config, theme string) error {
	availableThemes, err := getAvailableThemes(c.ThemePath)
	if err != nil {
		return fmt.Errorf("Error fetching available themes: %w\n", err)
	}
	if !slices.Contains(availableThemes, theme) {
		return errors.New("Invalid theme name.")
	}
	err = copy.CopyFile(filepath.Join(c.ThemePath, theme+".css"), filepath.Join(c.TargetPath, "assets", "index.css"))
	if err != nil {
		return fmt.Errorf("Error copying theme: %w\n", err)
	}
	err = config.SetConfig(c, "theme", theme+".css")
	if err != nil {
		return err
	}
	return nil
}

func getAvailableThemes(themesPath string) ([]string, error) {
	_, err := os.Stat(themesPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []string{}, fmt.Errorf("Theme directory '%s' does not exist.", themesPath)
		}
		return []string{}, err
	}
	var themes []string
	directoryContents, err := os.ReadDir(themesPath)
	for _, item := range directoryContents {
		itemName := item.Name()
		ext := filepath.Ext(filepath.Join(themesPath, itemName))
		if ext != ".css" {
			continue
		}
		themes = append(themes, itemName[:len(item.Name())-len(ext)])
	}
	return themes, nil
}
