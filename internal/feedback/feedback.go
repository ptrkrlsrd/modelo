package feedback

import "github.com/AlecAivazis/survey/v2"

type Answer struct {
	Name     string
	Template string
}

func AskForProjectName(answer *Answer) error {
	var firstQuestion = []*survey.Question{
		{
			Name:     "Name",
			Validate: survey.Required,
			Prompt: &survey.Input{
				Message: "Choose a name for your new project:",
			},
		},
	}

	return survey.Ask(firstQuestion, answer)
}

func AskForProjectTemplate(answer *Answer, options []string) error {
	var secondQuestion = []*survey.Question{
		{
			Name:     "Template",
			Validate: survey.Required,
			Prompt: &survey.Select{
				Message: "Choose a template:",
				Options: options,
			},
		},
	}

	return survey.Ask(secondQuestion, answer)
}
