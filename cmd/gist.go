package cmd

import (
	"context"
	"log"

	"github.com/ptrkrlsrd/modelo/internal/feedback"
	"github.com/ptrkrlsrd/modelo/pkg/github"
	"github.com/spf13/cobra"
)

var gistCmd = &cobra.Command{
	Use:   "gist",
	Short: "Gist",
	Run: func(cmd *cobra.Command, args []string) {
		service := github.NewService(config.GetString("username"), config.GetString("token"))
		ctx := context.Background()
		var selectedOption = new(feedback.Answer)
		selectedOption.Template = templateName
		selectedOption.FileName = gistFileName

		gists, err := service.GetGists(ctx)
		if err != nil {
			log.Fatalf("error getting gists: %s", err)
		}

		gistFiles := gists.CreateFileMap()
		gistNames := gists.GetFilenames()

		if err := feedback.AskGistQuestions("Select a Gist", selectedOption, gistNames); err != nil {
			log.Fatal(err)
		}

		selectedGist := gistFiles[selectedOption.Template]
		selectedGist.Write(selectedOption.ProjectName, selectedOption.FileName)
	},
}
