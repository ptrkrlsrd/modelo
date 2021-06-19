package github

type File struct {
	Name      string
	Extension string
	Text      string
}

type Gist struct {
	Name  string
	Files []File
}

type Gists []Gist
