package build

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/ajr-cabbage/sssg/src/internal/config"
	cp "github.com/nmrshll/go-cp"
	"github.com/yuin/goldmark"
)

func Build(c config.Config) error {
	ok := config.VerifyConfig(c)
	if !ok {
		return errors.New("Invalid or missing configuration\n'sssg init' to initialize a project in the current directory\n")
	}
	htmlTemplate, err := os.ReadFile(filepath.Join(c.TemplatePath, "default.html"))
	if err != nil {
		return err
	}
	err = generateDirectory(c.SourcePath, filepath.Join(c.TargetPath), string(htmlTemplate))
	err = copyAssets(c.AssetPath, c.TargetPath+"/assets/")
	if err != nil {
		return err
	}
	return nil
}

func copyAssets(sourceDirectory, targetDirectory string) error {
	_, err := os.Stat(targetDirectory)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = os.MkdirAll(targetDirectory, 0744)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	sourceDirContents, err := os.ReadDir(sourceDirectory)
	if err != nil {
		return err
	}
	for _, item := range sourceDirContents {
		itemSourcePath := filepath.Join(sourceDirectory, item.Name())
		itemTargetPath := filepath.Join(targetDirectory, item.Name())
		if item.IsDir() {
			copyAssets(itemSourcePath, itemTargetPath)
		} else {
			err := cp.CopyFile(itemSourcePath, itemTargetPath)
			if err != nil {
				fmt.Printf("Copy file error: %s\n", err.Error())
			}
		}
	}
	return nil
}

func generateDirectory(sourceDirectory, targetDirectory, template string) error {
	_, err := os.Stat(targetDirectory)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = os.MkdirAll(targetDirectory, 0744)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	sourceDirContents, err := os.ReadDir(sourceDirectory)
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	for _, item := range sourceDirContents {
		wg.Go(func() {
			itemSourcePath := filepath.Join(sourceDirectory, item.Name())
			itemTargetPath := filepath.Join(targetDirectory, item.Name())
			if item.IsDir() {
				err = generateDirectory(itemSourcePath, itemTargetPath, template)
				if err != nil {
					fmt.Println(err)
					return
				}
			} else {
				ext := filepath.Ext(itemSourcePath)
				if strings.ToLower(ext) != ".md" {
					fmt.Printf("Invalid file ext {%s} for: %s", ext, itemSourcePath)
					return
				}
				err := generatePage(itemSourcePath, itemTargetPath[:len(itemTargetPath)-len(ext)], template)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		})
	}
	wg.Wait()
	return nil
}

func generatePage(contentPath, targetPath, template string) error {
	md, err := os.ReadFile(contentPath)
	if err != nil {
		return err
	}
	html, err := mdToHTML(md)
	if err != nil {
		return err
	}
	htmlDoc := strings.Replace(template, "{{ Content }}", html, 1)
	err = os.WriteFile(targetPath+".html", []byte(htmlDoc), 0644)
	if err != nil {
		return err
	}
	return nil
}

func mdToHTML(md []byte) (string, error) {
	var buf bytes.Buffer
	err := goldmark.Convert(md, &buf)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// TODO
func getTitle(md string) (string, error) {
	return "", nil
}
