package github

import (
	"context"
	"os"
	"path"

	"github.com/shurcooL/githubv4"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

type Auth struct {
	Username string
	Token    string
}

type Service struct {
	GithubClient *githubv4.Client
	auth         Auth
}

func NewService(githubUsername string, githubToken string) Service {
	return Service{
		GithubClient: newClient(githubToken),
		auth: Auth{
			Username: githubUsername,
			Token:    githubToken,
		},
	}
}

type RepositoryQuery struct {
	Viewer struct {
		Name         string
		Repositories struct {
			Nodes Repositories
		} `graphql:"repositories(first: 100)"`
	}
}

func (service Service) GetRepositories(ctx context.Context) (Repositories, error) {
	var query RepositoryQuery
	if err := service.GithubClient.Query(ctx, &query, nil); err != nil {
		return nil, err
	}

	return query.Viewer.Repositories.Nodes, nil
}

type GistQuery struct {
	Viewer struct {
		Name  string
		Gists struct {
			Nodes Gists
		} `graphql:"gists(first: 100)"`
	}
}

func (service Service) GetGists(ctx context.Context) (Gists, error) {
	var query GistQuery
	if err := service.GithubClient.Query(ctx, &query, nil); err != nil {
		return nil, err
	}

	return query.Viewer.Gists.Nodes, nil
}

func (service Service) CloneTemplate(projectName string, template string, repositories Repositories) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	selectedRepo, err := repositories.FindRepoByName(template)
	if err != nil {
		return err
	}

	_, err = git.PlainClone(path.Join(dir, projectName), false, &git.CloneOptions{
		URL:      selectedRepo.URL,
		Progress: os.Stdout,
		Auth: &http.BasicAuth{
			Username: service.auth.Username,
			Password: service.auth.Token,
		},
	})

	if err != nil {
		return err
	}

	return nil
}
