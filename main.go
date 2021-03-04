package main

import (
	"context"
	"log"
	"os"

	"github.com/ptrkrlsrd/modelo/internal/feedback"
	"github.com/ptrkrlsrd/modelo/pkg/template"
)

func main() {
	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
		log.Println("No token set. Get one here: https://github.com/settings/tokens and set the $GITHUB_TOKEN environment variable")
		return
	}

	var selectedOption = new(feedback.Answer)
	err := feedback.AskForProjectName(selectedOption)
	if err != nil {
		log.Fatalf("error setting name: %s", err)
	}

	service := template.NewService(githubToken)
	repositories, err := service.GetRepositories(context.Background())
	if err != nil {
		log.Fatalf("error getting repositories: %s", err)
	}

	templateNames := repositories.GetTemplates().GetNames()
	err = feedback.AskForProjectTemplate(selectedOption, templateNames)
	if err != nil {
		log.Fatalf("error selecting repo: %s", err)
	}

	err = service.CloneTemplate(selectedOption.Name, selectedOption.Template, repositories)
	if err != nil {
		log.Fatalf("error cloning repo: %s", err)
	}
}
