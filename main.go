package main

import (
	"context"
	"log"
	"github.com/ptrkrlsrd/modelo/pkg/core"
	"os"

	"github.com/AlecAivazis/survey/v2"
)

type SurveyResponse struct {
	Name     string
	Template string
}

func askForProjectName(answer *SurveyResponse) error {
	var firstQ = []*survey.Question{
		{
			Name:     "Name",
			Validate: survey.Required,
			Prompt: &survey.Input{
				Message: "Choose a name for your new project:",
			},
		},
	}

	return survey.Ask(firstQ, answer)
}

func askForProjectTemplate(answers *SurveyResponse, options []string) error {
	var secondQ = []*survey.Question{
		{
			Name:     "Template",
			Validate: survey.Required,
			Prompt: &survey.Select{
				Message: "Choose a template:",
				Options: options,
			},
		},
	}

	return survey.Ask(secondQ, answers)
}

func main() {
	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
		log.Println("No token. Get one here: https://github.com/settings/tokens and set the $GITHUB_TOKEN environment variable")
		return
	}

	var selectedOption = new(SurveyResponse)
	err := askForProjectName(selectedOption)
	if err != nil {
		log.Fatalf("error setting name: %s", err)
	}

	service := core.NewService(githubToken)
	repositories, err := service.GetRepositories(context.Background())
	if err != nil {
		log.Fatalf("error getting repositories: %s", err)
	}

	templateNames := repositories.GetTemplates().GetNames()
	err = askForProjectTemplate(selectedOption, templateNames)
	if err != nil {
		log.Fatalf("error selecting repo: %s", err)
	}

	err = service.CloneTemplate(selectedOption.Name, selectedOption.Template, repositories)
	if err != nil {
		log.Fatalf("error cloning repo: %s", err)
	}
}
