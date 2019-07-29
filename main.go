package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/briandowns/spinner"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
	git "gopkg.in/src-d/go-git.v4"
)

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

func (repositories Repositories) GetTemplates() (templateRepositories Repositories) {
	for _, v := range repositories {
		if !v.IsTemplate || v.IsPrivate {
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

func getRepositories(client *githubv4.Client) (Repositories, error) {
	s := spinner.New(spinner.CharSets[10], 100*time.Millisecond)
	s.Start()
	defer s.Stop()

	var query GithubRepositoryQuery
	if err := client.Query(context.Background(), &query, nil); err != nil {
		return nil, err
	}

	return query.Viewer.Repositories.Nodes, nil
}

type Answers struct {
	Name     string
	Template string
}

func askFirstQuestion(answers *Answers) error {
	var firstQ = []*survey.Question{
		{
			Name:     "Name",
			Validate: survey.Required,
			Prompt: &survey.Input{
				Message: "Choose a name for your new project:",
			},
		},
	}

	return survey.Ask(firstQ, answers)
}

func askSecondQuestion(answers *Answers, repositories Repositories) error {
	templates := repositories.GetTemplates()
	templateNames := templates.GetNames()

	var secondQ = []*survey.Question{
		{
			Name:     "Template",
			Validate: survey.Required,
			Prompt: &survey.Select{
				Message: "Choose a template:",
				Options: templateNames,
			},
		},
	}

	return survey.Ask(secondQ, answers)
}

func cloneRepo(answers *Answers, repositories Repositories) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	selectedRepo, err := repositories.FindRepoByName(answers.Template)
	if err != nil {
		return err
	}

	_, err = git.PlainClone(fmt.Sprintf("%s/%s", dir, answers.Name), false, &git.CloneOptions{
		URL:      selectedRepo.URL,
		Progress: os.Stdout,
	})

	if err != nil {
		return err
	}

	return nil
}

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Println("No token. Get one here: https://github.com/settings/tokens and set the $GITHUB_TOKEN environment variable")
		return
	}

	var answers = new(Answers)
	err := askFirstQuestion(answers)
	if err != nil {
		log.Fatalf("error setting name: %s", err)
	}

	client := newClient(token)
	repositories, err := getRepositories(client)
	if err != nil {
		log.Fatalf("error getting repositories: %s", err)
	}

	err = askSecondQuestion(answers, repositories)
	if err != nil {
		log.Fatalf("error selecting repo: %s", err)
	}

	err = cloneRepo(answers, repositories)
	if err != nil {
		log.Fatalf("error cloning repo: %s", err)
	}
}
