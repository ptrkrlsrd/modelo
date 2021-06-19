package cmd

import (
	"context"
	"log"
	"os"
	"path"

	"github.com/ptrkrlsrd/modelo/internal/feedback"
	"github.com/ptrkrlsrd/modelo/pkg/github"
	"github.com/spf13/viper"
)

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

func Execute() {
	config, err := readConfig()
	if err != nil {
		log.Fatal(err)
	}

	githubToken := config.GetString("token")
	githubUsername := config.GetString("username")

	service := github.NewService(githubUsername, githubToken)
	ctx := context.Background()
	var selectedOption = new(feedback.TemplateAnswer)
	if err = feedback.AskForProjectName(selectedOption); err != nil {
		log.Fatal(err)
	}

	args := os.Args
	if len(args) == 1 {
		selectFromGithubTemplates(service, ctx, selectedOption)
	} else if args[1] == "from-gist" {
		selectFromGithubGists(service, ctx, selectedOption)
	}

}

func selectFromGithubGists(service github.Service, ctx context.Context, selectedOption *feedback.TemplateAnswer) {
	gists, err := service.GetGists(ctx)
	if err != nil {
		log.Fatalf("error getting gists: %s", err)
	}

	gistFiles, gistNames := extractFilesFromGists(gists)
	if err = feedback.AskForProjectGist(selectedOption, gistNames); err != nil {
		log.Fatalf("error selecting repo: %s", err)
	}

	if err = os.Mkdir(selectedOption.ProjectName, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	gist := gistFiles[selectedOption.Template]
	writeGistToFile(selectedOption.ProjectName, gist)
}

func selectFromGithubTemplates(service github.Service, ctx context.Context, selectedOption *feedback.TemplateAnswer) {
	repositories, err := service.GetRepositories(ctx)
	if err != nil {
		log.Fatalf("error getting repositories: %s", err)
	}

	templates := repositories.FilterTemplates()
	templateNames := templates.Names()
	if err = feedback.AskForProjectTemplate(selectedOption, templateNames); err != nil {
		log.Fatalf("error selecting repo: %s", err)
	}

	if err = service.CloneTemplate(selectedOption.ProjectName, selectedOption.Template, repositories); err != nil {
		log.Fatalf("error cloning repo: %s", err)
	}
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
