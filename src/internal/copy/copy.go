package copy

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(src, dest string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("Cannot open file '%s': %w", src, err)
	}
	defer srcFile.Close()
	destFile, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("Cannot create file '%s': %w", dest, err)
	}
	defer destFile.Close()
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return fmt.Errorf("Error copying '%s' to '%s': %w", src, dest, err)
	}
	return destFile.Close()
}
