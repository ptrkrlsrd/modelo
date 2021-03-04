package feedback

import "github.com/AlecAivazis/survey/v2"

type Answer struct {
	Name     string
	Template string
}

func AskForProjectName(answer *Answer) error {
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

func AskForProjectTemplate(answer *Answer, options []string) error {
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

	return survey.Ask(secondQ, answer)
}
