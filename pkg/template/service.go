package template

import (
	"context"
	"fmt"
	"os"

	"github.com/shurcooL/githubv4"
	"gopkg.in/src-d/go-git.v4"
)

type Service struct {
	GithubClient *githubv4.Client
}

func NewService(githubToken string) Service {
	return Service{GithubClient: newGithubClient(githubToken)}
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
	})

	if err != nil {
		return err
	}

	return nil
}
