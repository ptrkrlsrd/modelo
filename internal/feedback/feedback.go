package feedback

import (
	"github.com/AlecAivazis/survey/v2"
)

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
	var questions []*survey.Question

	if answer.Template == "" {
		questions = append(questions, &survey.Question{
			Name:     "Template",
			Validate: survey.Required,
			Prompt: &survey.Select{
				Message: message,
				Options: options,
				VimMode: true,
			},
		})
	}

	if answer.ProjectName == "" {
		questions = append(questions, &projectNameQuestion)
	}

	return survey.Ask(questions, answer)
}

func AskGistQuestions(message string, answer *Answer, options []string) error {
	var questions []*survey.Question

	if answer.Template == "" {
		questions = append(questions, &survey.Question{
			Name:     "Template",
			Validate: survey.Required,
			Prompt: &survey.Select{
				Renderer: survey.Renderer{},
				Message:  message,
				Options:  options,
				VimMode:  true,
			},
		})
	}

	if answer.FileName == "" {
		questions = append(questions, &survey.Question{
			Name: "FileName",
			Prompt: &survey.Input{
				Renderer: survey.Renderer{},
				Help:     "Leave blank to use Gist name",
				Message:  "Filename: ",
			},
		})
	}

	return survey.Ask(questions, answer)
}
