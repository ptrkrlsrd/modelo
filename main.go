package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ptrkrlsrd/modelo/internal/feedback"
	"github.com/ptrkrlsrd/modelo/pkg/template"
	"github.com/spf13/viper"
)

func readConfig() (*viper.Viper, error) {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")

	configPaths := []string{"/etc/modelo/", "$HOME/.config/modelo", ".modelo"}
	for _, i := range configPaths {
		config.AddConfigPath(i)
	}

	if err := config.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("fatal error config file: %s", err)
	}

	return config, nil
}

func main() {
	config, err := readConfig()
	if err != nil {
		log.Fatalf("error parsing config: %s", err)
	}

	githubToken := config.Get("token").(string)
	githubUsername := config.Get("username").(string)

	var selectedOption = new(feedback.Answer)
	if err = feedback.AskForProjectName(selectedOption); err != nil {
		log.Fatalf("error setting name: %s", err)
	}

	service := template.NewService(githubUsername, githubToken)
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
