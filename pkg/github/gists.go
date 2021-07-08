package github

type Gist struct {
	Name  string
	Files Files
}

type Gists []Gist

func (gists Gists) GetFiles() Files {
	var files Files
	for _, i := range gists {
		files = append(files, i.Files...)
	}
	return files
}
