package build

import (
	"bytes"
	"os"
	"strings"

	"github.com/yuin/goldmark"
)

type Config struct {
	SourcePath   string
	TargetPath   string
	AssetPath    string
	TemplatePath string
}

func (c *Config) Build() error {
	// copy assets to target directory
	// call a recursive directory walk function
	return nil
}

func generateDirectory(sourceDirectory, targetDirectory, template string) error {
	// this function can make good use of go funcs
	// both recursive calls and calls to GeneratePage can run without mutexes
	// recursively walk the source directory
	// base case is finding an actual file, checking if it's MD and calling GeneratePage
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
	err = os.WriteFile(targetPath, []byte(htmlDoc), 0644)
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

func getTitle(md string) (string, error) {
	return "", nil
}
