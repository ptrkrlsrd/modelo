package feedback

import "github.com/AlecAivazis/survey/v2"

type Answer struct {
	ProjectName string
	FileName    string
	Template    string
}

func newQuestion(name string, message string, answer *Answer) survey.Question {
	var firstQuestion = survey.Question{
		Name:     name,
		Validate: survey.Required,
		Prompt: &survey.Input{
			Message: message,
		},
	}

	return firstQuestion
}

func AskTemplateQuestion(message string, answer *Answer, options []string) error {
	projectNameQuestion := newQuestion("ProjectName", "Choose a name for your new project:", answer)

	var questions = []*survey.Question{
		&projectNameQuestion,
		{
			Name:     "Template",
			Validate: survey.Required,
			Prompt: &survey.Select{
				Message: message,
				Options: options,
			},
		},
	}

	return survey.Ask(questions, answer)
}

func AskGistQuestions(message string, answer *Answer, options []string) error {
	fileNameQuestion := newQuestion("FileName", "Filename", answer)

	var questions = []*survey.Question{
		&fileNameQuestion,
		{
			Name:     "Template",
			Validate: survey.Required,
			Prompt: &survey.Select{
				Message: message,
				Options: options,
			},
		},
	}

	return survey.Ask(questions, answer)
}
