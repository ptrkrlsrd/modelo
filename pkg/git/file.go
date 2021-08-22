package git

import (
	"fmt"
	"os"
	"path"
)

type File struct {
	Name      string
	Extension string
	Text      string
}

type Files []File

func (file *File) Write(filePath string, fileName string) error {
	gistFile, err := os.Create(path.Join(filePath, fileName))
	if err != nil {
		return err
	}

	defer gistFile.Close()
	os.Mkdir(filePath, os.ModePerm)

	if _, err = gistFile.WriteString(file.Text); err != nil {
		return err
	}

	return nil
}

func contains(item string, slice []string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func (files Files) Filter(ignored []string) Files {
	var filteredFiles Files
	for _, file := range files {
		if contains(file.Name, ignored) {
			continue
		}
		filteredFiles = append(filteredFiles, file)
	}
	return filteredFiles
}

func (files Files) GetNames() []string {
	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name)
	}
	return fileNames
}

func (files Files) ToMap() (map[string]File, error) {
	fileMap := make(map[string]File)

	for _, file := range files {
		if file.Name == "" {
			return nil, fmt.Errorf("filename cannot be empty")
		}
		fileMap[file.Name] = file
	}
	return fileMap, nil
}
