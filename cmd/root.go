package cmd

import (
	"context"
	"log"
	"os"

	"github.com/ptrkrlsrd/modelo/internal/feedback"
	"github.com/ptrkrlsrd/modelo/pkg/git"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	config         *viper.Viper
	selectedOption *feedback.Answer
	projectPath    string
	templateName   string
	gistFileName   string
)

func init() {
	var err error
	config, err = readConfig()
	if err != nil {
		log.Fatal(err)
	}

	rootCmd.PersistentFlags().StringVarP(&templateName, "template", "t", "", "template name")
	rootCmd.PersistentFlags().StringVarP(&projectPath, "path", "p", "", "path")
	gistCmd.PersistentFlags().StringVarP(&gistFileName, "filename", "f", "", "file name")

	rootCmd.AddCommand(gistCmd)
}

var rootCmd = &cobra.Command{
	Use:   "modelo",
	Short: "Boilerplate your projects from Github Templates and Gists",
	Run: func(cmd *cobra.Command, args []string) {
		selectedOption = &feedback.Answer{
			Template:    templateName,
			ProjectName: projectPath,
		}

		service := git.NewService(config.GetString("username"), config.GetString("token"))
		ctx := context.Background()

		ignored := config.GetStringSlice("repositories.ignored")
		repositories, err := service.GetRepositories(ctx)
		filteredRepos := repositories.Filter(ignored)
		if err != nil {
			log.Fatalf("error getting repositories: %s", err)
		}

		includedRepos := config.GetStringMapString("repositories.include")
		for i, v := range includedRepos {
			filteredRepos = filteredRepos.AddRepository(git.NewRepository(i, v, false, true))
		}

		templates := filteredRepos.GetTemplates()
		templateNames := templates.GetNames()

		if err = feedback.AskTemplateQuestion("Select a Github Template: ", selectedOption, templateNames); err != nil {
			log.Fatalf("error selecting repo: %s", err)
		}

		if err = service.CloneTemplate(selectedOption.ProjectName, selectedOption.Template, filteredRepos); err != nil {
			log.Fatalf("error cloning repo: %s", err)
		}
	},
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
