package death

import (
	"os"
	"path/filepath"

	"github.com/qean32/holo/internal/constants"
	errorlist "github.com/qean32/holo/internal/constants/error-list"
)

func writeToFile(path string, flag int, data string, errFunc func(string, error) error) error {
	file, err := os.OpenFile(filepath.Join(constants.Root, path), flag, 0o644)
	if err != nil {
		return errFunc(path, err)
	}
	defer file.Close()

	if _, err := file.WriteString(data); err != nil {
		return errFunc(path, err)
	}
	return nil
}

func WriteFile(data, path string) error {
	return writeToFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, data, errorlist.CreateFile)
}

func PushToFile(path, newText string) error {
	return writeToFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, newText+"\n", errorlist.WriteFile)
}

func CreateFile(path string, content string) error {
	if err := os.WriteFile(filepath.Join(constants.Root, path), []byte(content), 0o644); err != nil {
		return errorlist.CreateFile(path, err)
	}
	return nil
}

func ClearFile(path string) error {
	return WriteFile("", path)
}

func RemoveFile(path string) error {
	if err := os.Remove(filepath.Join(constants.Root, path)); err != nil {
		return errorlist.RemoveFile(path, err)
	}
	return nil
}

func ListFilesByExt(dir string, ext string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, errorlist.ReadFile(dir, err)
	}

	var files []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if filepath.Ext(entry.Name()) == ext {
			files = append(files, entry.Name())
		}
	}
	return files, nil
}
