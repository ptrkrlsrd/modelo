package template

import (
	"context"
	"fmt"
	"os"

	"github.com/shurcooL/githubv4"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

type GithubAuth struct {
	Username string
	Token    string
}

type Service struct {
	GithubClient *githubv4.Client
	auth         GithubAuth
}

func NewService(githubUsername string, githubToken string) Service {
	return Service{
		GithubClient: newGithubClient(githubToken),
		auth: GithubAuth{
			Username: githubUsername,
			Token:    githubToken,
		},
	}
}

type GithubRepositoryQuery struct {
	Viewer struct {
		Name         string
		Repositories struct {
			Nodes Repositories
		} `graphql:"repositories(first: 100)"`
	}
}

func (service Service) GetRepositories(ctx context.Context) (Repositories, error) {
	var query GithubRepositoryQuery
	if err := service.GithubClient.Query(ctx, &query, nil); err != nil {
		return nil, err
	}

	return query.Viewer.Repositories.Nodes, nil
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

	_, err = git.PlainClone(fmt.Sprintf("%s/%s", dir, projectName), false, &git.CloneOptions{
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
