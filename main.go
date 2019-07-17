package main

import (
	"context"
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

func main() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)
	var q struct {
		Viewer struct {
			Name         string
			Repositories struct {
				Nodes []Repository
			} `graphql:"repositories(first: 100)"`
		}
	}

	err := client.Query(context.Background(), &q, nil)
	if err != nil {
		log.Fatal(err)
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	var templates []string
	var repositories = q.Viewer.Repositories.Nodes
	for _, v := range q.Viewer.Repositories.Nodes {
		if v.IsTemplate && !v.IsPrivate {
			templates = append(templates, v.Name)
		}
	}

	answers := struct {
		Name     string
		Template string
	}{}

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

	err = survey.Ask(qs, &answers)
	if err != nil {
		log.Println(err.Error())
		return
	}

	for _, v := range repositories {
		if v.Name == answers.Template {
			root := dir + "/" + answers.Name
			_, err := git.PlainClone(root, false, &git.CloneOptions{
				URL:      v.URL,
				Progress: os.Stdout,
			})

			if err != nil {
				log.Println(err.Error())
				return
			}
		}
	}
}
