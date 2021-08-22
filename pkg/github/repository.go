package github

import (
	"errors"
	"fmt"
)

type Repository struct {
	Name       string
	URL        string
	IsTemplate bool
	IsPrivate  bool
}

type Repositories []Repository

func (repositories Repositories) GetTemplates() Repositories {
	var templateRepositories Repositories
	for _, v := range repositories {
		if !v.IsTemplate {
			continue
		}

		templateRepositories = append(templateRepositories, v)
	}

	return templateRepositories
}

func (repositories Repositories) GetNames() (repositoryNames []string) {
	for _, v := range repositories {
		repositoryNames = append(repositoryNames, v.Name)
	}

	return repositoryNames
}

func indexOf(element string, ignored []string) (int, error) {
	for k, v := range ignored {
		if element == v {
			return k, nil
		}
	}
	return -1, errors.New("could not find")
}

func (repositories Repositories) Filter(ignored []string) Repositories {
	filtered := Repositories{}
	for _, v := range repositories {
		if _, err := indexOf(v.Name, ignored); err != nil {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

func (repositories Repositories) FindByName(name string) (Repository, error) {
	for _, v := range repositories {
		if v.Name != name {
			continue
		}

		return v, nil
	}

	return Repository{}, fmt.Errorf("failed finding repository with name: " + name)
}
