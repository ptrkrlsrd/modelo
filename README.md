# Modelo
## Boilerplate your projects from Github Templates and Gists 
[![Quality gate](https://sonarcloud.io/api/project_badges/quality_gate?project=ptrkrlsrd_modelo)](https://sonarcloud.io/dashboard?id=ptrkrlsrd_modelo)

## Intro
Project templating made easy using Github templates or create single files from Gists. Think of it as create-react-app for any type of projects!

### Demo
#### Create a project from a Gihub Template repository
![](recorded.gif)


### Usage
```
Boilerplate your projects from Github Templates and Gists

Usage:
  modelo [flags]
  modelo [command]

Available Commands:
  gist        Gist
  help        Help about any command

Flags:
  -h, --help              help for modelo
  -p, --path string       path
  -t, --template string   template name

Use "modelo [command] --help" for more information about a command.
```

1. Create a file called 'config.json' in the following path: $HOME/.config/modelo
2. Create a personal access token on Github with read access to repositories
3. Add the following content to 'config.json': 
``` json
        { "username": "<github username>", "token": "<github token>" } 
```
4. Run `modelo` or `modelo gist` (to create from a Gist) and follow the instructions
* Note: Only 'template repositories' will be listed in the CLI

## Optional
* If you want to ignore a repo, you can do so by adding the name of the repo to `repos.ignored` as shown below. You also ignore gists by adding the name of the gist to `gists.ignored`.
``` json
{ 
  "username": "<github username>", 
  "token": "<github token>",
  "repos": {
    "ignored": ["<your ignored repo>"],
  },
  "gists": {
    "ignored": [
      "main.py", 
    ]
  }
} 
```

## Upcoming features
* Add ability to use other Github users templates

## What does Modelo mean?
Modelo is Portuguese for template. I chose the name because I have a good friend from Portugal, who recently lost a loved one due to cancer.
