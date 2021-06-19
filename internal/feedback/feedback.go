package feedback

import "github.com/AlecAivazis/survey/v2"

type TemplateAnswer struct {
	ProjectName string
	Template    string
}

func AskForProjectName(answer *TemplateAnswer) error {
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

func AskForProjectTemplate(answer *TemplateAnswer, options []string) error {
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

func AskForProjectGist(answer *TemplateAnswer, options []string) error {
	var secondQuestion = []*survey.Question{
		{
			Name:     "Template",
			Validate: survey.Required,
			Prompt: &survey.Select{
				Message: "Choose a gist:",
				Options: options,
			},
		},
	}

	return survey.Ask(secondQuestion, answer)
}
