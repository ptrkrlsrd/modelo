package github

import (
	"os"
	"path"
)

type File struct {
	Name      string
	Extension string
	Text      string
}

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

type Gist struct {
	Name  string
	Files []File
}

type Gists []Gist

func (gists Gists) GetFilenames() []string {
	var gistNames []string
	for _, i := range gists {
		for _, file := range i.Files {
			gistNames = append(gistNames, file.Name)
		}
	}
	return gistNames
}

func (gists Gists) CreateFileMap() map[string]File {
	fileMap := make(map[string]File)

	for _, i := range gists {
		for _, file := range i.Files {
			fileMap[file.Name] = file
		}
	}
	return fileMap
}
