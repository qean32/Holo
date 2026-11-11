package death

import (
	"os"
	"path/filepath"

	"main/constants"
	errorlist "main/constants/error-list"
)

func WriteFile(data, path string) error {
	file, err := os.Create(filepath.Join(constants.Root, path))
	if err != nil {
		return errorlist.CreateFile(path, err)
	}
	defer file.Close()

	if _, err := file.WriteString(data); err != nil {
		return errorlist.WriteFile(path, err)
	}
	return nil
}

func PushToFile(path, newText string) error {
	file, err := os.OpenFile(filepath.Join(constants.Root, path), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return errorlist.OpenFile(path, err)
	}
	defer file.Close()

	if _, err := file.WriteString(newText + "\n"); err != nil {
		return errorlist.WriteFile(path, err)
	}
	return nil
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
