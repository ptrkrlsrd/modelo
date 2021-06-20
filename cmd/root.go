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

var config *viper.Viper

var rootCmd = &cobra.Command{
	Use:   "modelo",
	Short: "Boilerplate your projects from Github Templates and Gists",
	Run: func(cmd *cobra.Command, args []string) {
		service := github.NewService(config.GetString("username"), config.GetString("token"))
		ctx := context.Background()
		selectFromGithubTemplates(service, ctx)
	},
}

func init() {
	var err error
	config, err = readConfig()
	if err != nil {
		log.Fatal(err)
	}

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

func selectFromGithubTemplates(service github.Service, ctx context.Context) {
	repositories, err := service.GetRepositories(ctx)
	if err != nil {
		log.Fatalf("error getting repositories: %s", err)
	}

	templates := repositories.FilterTemplates()
	templateNames := templates.Names()

	var selectedOption = new(feedback.Answer)
	if err = feedback.AskTemplateQuestion("Select a Github Template: ", selectedOption, templateNames); err != nil {
		log.Fatalf("error selecting repo: %s", err)
	}

	if err = service.CloneTemplate(selectedOption.ProjectName, selectedOption.Template, repositories); err != nil {
		log.Fatalf("error cloning repo: %s", err)
	}
}
