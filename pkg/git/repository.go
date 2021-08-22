package git

import (
	"fmt"
	"regexp"
)

type Repository struct {
	Name       string
	URL        string
	IsTemplate bool
	IsPrivate  bool
}

func NewRepository(name, url string, isPrivate, isTemplate bool) Repository {
	return Repository{
		Name:       name,
		URL:        url,
		IsTemplate: isTemplate,
		IsPrivate:  isPrivate,
	}
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

func (repositories Repositories) Filter(ignored []string) Repositories {
	filtered := Repositories{}
	for _, v := range repositories {
		if !contains(v.Name, ignored) {
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

func (repositories Repositories) AddRepository(repo Repository) Repositories {
	repositories = append(repositories, repo)
	return repositories
}

func IsValidGitURL(url string) bool {
	validGitURL := regexp.MustCompile(`((git|ssh|http(s)?)|(git@[\w\.]+))(:(//)?)([\w\.@\:/\-~]+)(\.git)?(/)?`)
	return validGitURL.Match([]byte(url))
}

func IsValidRepoName(name string) bool {
	validGitName := regexp.MustCompile(`[A-Za-z0-9_.-]`)
	return validGitName.Match([]byte(name))
}
