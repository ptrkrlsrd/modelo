package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
	"gopkg.in/src-d/go-git.v4"
)

var query struct {
	Viewer struct {
		Login          githubv4.String
		CreatedAt      githubv4.DateTime
		IsBountyHunter githubv4.Boolean
		BioHTML        githubv4.HTML
		WebsiteURL     githubv4.URI
	}
}

type Repository struct {
	Name       string
	URL        string
	IsTemplate bool
	IsPrivate  bool
}

type Repositories []Repository

type GithubRepositoryQuery struct {
	Viewer struct {
		Name         string
		Repositories struct {
			Nodes Repositories
		} `graphql:"repositories(first: 100)"`
	}
}

func (repositories Repositories) GetNames() (templates []string) {
	for _, v := range repositories {
		if !v.IsTemplate || v.IsPrivate {
			continue
		}

		templates = append(templates, v.Name)
	}
	return templates
}

func (repositories Repositories) FindRepoByName(name string) (Repository, error) {
	for _, v := range repositories {
		if v.Name != name {
			continue
		}

		return v, nil
	}

	return Repository{}, fmt.Errorf("failed finding repository with name: " + name)
}

func newClient(token string) *githubv4.Client {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	return githubv4.NewClient(httpClient)
}

func main() {
	client := newClient(os.Getenv("GITHUB_TOKEN"))

	var query GithubRepositoryQuery
	if err := client.Query(context.Background(), &query, nil); err != nil {
		log.Fatal(err)
		return
	}

	var repositories = query.Viewer.Repositories.Nodes
	templates := repositories.GetNames()

	var qs = []*survey.Question{
		{
			Name:     "Name",
			Validate: survey.Required,
			Prompt: &survey.Input{
				Message: "Choose a name for your new project:",
			},
		},
		{
			Name:     "Template",
			Validate: survey.Required,
			Prompt: &survey.Select{
				Message: "Choose a template:",
				Options: templates,
			},
		},
	}

	answers := struct {
		Name     string
		Template string
	}{}

	if err := survey.Ask(qs, &answers); err != nil {
		log.Println(err.Error())
		return
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return
	}

	selectedRepo, err := repositories.FindRepoByName(answers.Name)
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = git.PlainClone(fmt.Sprintf("%s/%s", dir, answers.Name), false, &git.CloneOptions{
		URL:      selectedRepo.URL,
		Progress: os.Stdout,
	})

	if err != nil {
		log.Println(err.Error())
		return
	}
}
