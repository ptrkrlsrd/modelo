package feedback

import "github.com/AlecAivazis/survey/v2"

type Answer struct {
	ProjectName string
	Template    string
}

func AskForProjectName(answer *Answer) error {
	var firstQuestion = []*survey.Question{
		{
			Name:     "ProjectName",
			Validate: survey.Required,
			Prompt: &survey.Input{
				Message: "Choose a name for your new project:",
			},
		},
	}

	return survey.Ask(firstQuestion, answer)
}

func AskForTemplate(message string, answer *Answer, options []string) error {
	var secondQuestion = []*survey.Question{
		{
			Name:     "Template",
			Validate: survey.Required,
			Prompt: &survey.Select{
				Message: message,
				Options: options,
			},
		},
	}

	return survey.Ask(secondQuestion, answer)
}
