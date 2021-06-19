package cmd

import (
	"context"
	"log"
	"os"

	"github.com/ptrkrlsrd/modelo/internal/feedback"
	"github.com/ptrkrlsrd/modelo/pkg/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "modelo",
	Short: "Boilerplate your projects from Github Templates and gists",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := readConfig()
		if err != nil {
			log.Fatal(err)
		}

		githubToken := config.GetString("token")
		githubUsername := config.GetString("username")

		service := github.NewService(githubUsername, githubToken)
		ctx := context.Background()
		var selectedOption = new(feedback.Answer)
		if err = feedback.AskForProjectName(selectedOption); err != nil {
			log.Fatal(err)
		}

		selectFromGithubTemplates(service, ctx, selectedOption)
	},
}

func init() {
	rootCmd.AddCommand(gistCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func readConfig() (*viper.Viper, error) {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")

	config.AddConfigPath("$HOME/.config/modelo")

	if err := config.ReadInConfig(); err != nil {
		return nil, err
	}

	return config, nil
}

func selectFromGithubTemplates(service github.Service, ctx context.Context, selectedOption *feedback.Answer) {
	repositories, err := service.GetRepositories(ctx)
	if err != nil {
		log.Fatalf("error getting repositories: %s", err)
	}

	templates := repositories.FilterTemplates()
	templateNames := templates.Names()
	if err = feedback.AskForTemplate("Select a Github Template: ", selectedOption, templateNames); err != nil {
		log.Fatalf("error selecting repo: %s", err)
	}

	if err = service.CloneTemplate(selectedOption.ProjectName, selectedOption.Template, repositories); err != nil {
		log.Fatalf("error cloning repo: %s", err)
	}
}
