/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"log"

	"github.com/ptrkrlsrd/modelo/pkg/git"
	"github.com/spf13/cobra"
)

// repositoryCmd represents the repository command
var repositoryCmd = &cobra.Command{
	Use: "repository",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			cmd.Help()
			return nil
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var includeRepositoryCmd = &cobra.Command{
	Use:  "include",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		repoName := args[0]
		repoURL := args[1]

		if !git.IsValidGitURL(repoURL) {
			log.Fatal("invalid repo url")
		}

		if !git.IsValidRepoName(repoURL) {
			log.Fatal("invalid repo name")
		}

		includedRepos := config.GetStringMapString("repositories.include")
		includedRepos[repoName] = repoURL

		config.Set("repositories.include", includedRepos)
		config.WriteConfig()
	},
}

func init() {
	repositoryCmd.AddCommand(includeRepositoryCmd)
	rootCmd.AddCommand(repositoryCmd)
}
