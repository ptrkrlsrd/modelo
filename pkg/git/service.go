package git

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

type RepoReaderWriter interface {
	Query(ctx context.Context, q interface{}, variables map[string]interface{}) error
	Mutate(ctx context.Context, m interface{}, input githubv4.Input, variables map[string]interface{}) error
}

type Service struct {
	GithubClient RepoReaderWriter
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

type repositoryData struct {
	Nodes Repositories
}

type repositoryViewer struct {
	Name         string
	Repositories repositoryData `graphql:"repositories(first: 100)"`
}

type RepositoryQuery struct {
	Viewer repositoryViewer
}

func (service Service) GetRepositories(ctx context.Context) (Repositories, error) {
	var query RepositoryQuery
	if err := service.GithubClient.Query(ctx, &query, nil); err != nil {
		return nil, err
	}

	return query.Viewer.Repositories.Nodes, nil
}

type gistData struct {
	Nodes Gists
}

type gistViewer struct {
	Name  string
	Gists gistData `graphql:"gists(first: 100)"`
}

type GistQuery struct {
	Viewer gistViewer
}

func (service Service) GetUsersGists(ctx context.Context) (Gists, error) {
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

	selectedRepo, err := repositories.FindByName(template)
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
