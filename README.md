# Modelo
[![Quality gate](https://sonarcloud.io/api/project_badges/quality_gate?project=ptrkrlsrd_modelo)](https://sonarcloud.io/dashboard?id=ptrkrlsrd_modelo)
## Project templating made easy. 

![](recorded.gif)

## Intro
Project templating made easy using Github templates. Think of it as create-react-app for any type of projects.

## Usage
1. Create a file called 'config.json' in one of the following paths: [/etc/modelo/ $HOME/.config/modelo .modelo/]
2. Create a personal access token on Github with read access to repositories
3. Add the following content: 
``` json
        { "username": "<github username>", "token": "<github token>" } 
```
* Note: Only 'template repositories' will be listed in the CLI

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

## TODO
* Allow users to use other Github users templates

## What does Modelo mean?
Modelo is Portuguese for template. I chose the name because I have a good friend from Portugal, who recently lost a loved one due to cancer.
