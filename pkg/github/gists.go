package github

type Gist struct {
	Name  string
	Files Files
}

type Gists []Gist

func (gists Gists) GetFiles() Files {
	var files Files
	for _, i := range gists {
		for _, file := range i.Files {
			files = append(files, file)
		}
	}
	return files
}
