package cmd

import (
	"context"
	"log"

	"github.com/ptrkrlsrd/modelo/internal/feedback"
	"github.com/ptrkrlsrd/modelo/pkg/git"
	"github.com/spf13/cobra"
)

var gistCmd = &cobra.Command{
	Use:   "gist",
	Short: "Gist",
	Run: func(cmd *cobra.Command, args []string) {
		service := git.NewService(config.GetString("username"), config.GetString("token"))
		ctx := context.Background()
		selectedOption = &feedback.Answer{
			Template: templateName,
			FileName: gistFileName,
		}

		gists, err := service.GetGists(ctx)
		if err != nil {
			log.Fatalf("error getting gists: %s", err)
		}

		ignored := config.GetStringSlice("gists.ignored")
		filteredFiles := gists.GetFiles().Filter(ignored)
		gistNames := filteredFiles.GetNames()

		if err := feedback.AskGistQuestions("Select a Gist", selectedOption, gistNames); err != nil {
			log.Fatal(err)
		}

		if selectedOption.FileName == "" {
			selectedOption.FileName = selectedOption.Template
		}

		gistFiles, err := filteredFiles.ToMap()
		if err != nil {
			log.Fatal(err)
		}

		selectedGist := gistFiles[selectedOption.Template]
		selectedGist.Write(selectedOption.ProjectName, selectedOption.FileName)
	},
}
