package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ptrkrlsrd/modelo/internal/feedback"
	"github.com/ptrkrlsrd/modelo/pkg/github"
	"github.com/spf13/viper"
)

func readConfig() (*viper.Viper, error) {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")

	configPaths := []string{"/etc/modelo/", "$HOME/.config/modelo", ".modelo/"}
	for _, i := range configPaths {
		config.AddConfigPath(i)
	}

	if err := config.ReadInConfig(); err != nil {
		helpString := fmt.Sprintf("no config file named 'config.json'\n")
		helpString += fmt.Sprintf("1. Create a file called 'config.json' in one of the following paths: %s\n", configPaths)
		helpString += "2. Create a personal access token on Github with read access to repositories\n"
		helpString += fmt.Sprintf("3. Add the following content: \n\t%s\n", `{ "username": "<github username>", "token": "<github token>" } `)

		err = fmt.Errorf("fatal error config file: %s", helpString)
		return nil, err
	}

	return config, nil
}

func main() {
	config, err := readConfig()
	if err != nil {
		log.Fatal(err)
	}

	githubToken := config.Get("token").(string)
	githubUsername := config.Get("username").(string)

	var selectedOption = new(feedback.Answer)
	if err = feedback.AskForProjectName(selectedOption); err != nil {
		log.Fatalf("error setting name: %s", err)
	}

	service := github.NewService(githubUsername, githubToken)
	repositories, err := service.GetRepositories(context.Background())
	if err != nil {
		log.Fatalf("error getting repositories: %s", err)
	}

	templates := repositories.GetTemplates()
	templateNames := templates.GetNames()
	if err = feedback.AskForProjectTemplate(selectedOption, templateNames); err != nil {
		log.Fatalf("error selecting repo: %s", err)
	}

	if err = service.CloneTemplate(selectedOption.Name, selectedOption.Template, repositories); err != nil {
		log.Fatalf("error cloning repo: %s", err)
	}
}
