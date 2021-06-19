package cmd

import (
	"context"
	"log"
	"os"
	"path"

	"github.com/ptrkrlsrd/modelo/internal/feedback"
	"github.com/ptrkrlsrd/modelo/pkg/github"
	"github.com/spf13/cobra"
)

var gistCmd = &cobra.Command{
	Use:   "gist",
	Short: "Gist",
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
		gists, err := service.GetGists(ctx)
		if err != nil {
			log.Fatalf("error getting gists: %s", err)
		}

		gistFiles, gistNames := extractFilesFromGists(gists)
		if err = feedback.AskForTemplate("Select a Github Gist", selectedOption, gistNames); err != nil {
			log.Fatalf("error selecting repo: %s", err)
		}

		createFolder(selectedOption.ProjectName)

		gist := gistFiles[selectedOption.Template]
		writeGistToFile(selectedOption.ProjectName, gist)
	},
}

func createFolder(folderPath string) {
	os.Mkdir(folderPath, os.ModePerm)
}

func extractFilesFromGists(gists github.Gists) (g map[string]github.File, gistNames []string) {
	g = make(map[string]github.File)

	for _, i := range gists {
		for _, file := range i.Files {
			g[file.Name] = file
			gistNames = append(gistNames, file.Name)
		}
	}
	return g, gistNames
}

func writeGistToFile(filePath string, gist github.File) error {
	gistFile, err := os.Create(path.Join(filePath, gist.Name))
	if err != nil {
		return err
	}

	defer gistFile.Close()

	if _, err = gistFile.WriteString(gist.Text); err != nil {
		return err
	}

	return nil
}
