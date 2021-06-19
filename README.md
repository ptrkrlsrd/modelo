# Modelo
[![Quality gate](https://sonarcloud.io/api/project_badges/quality_gate?project=ptrkrlsrd_modelo)](https://sonarcloud.io/dashboard?id=ptrkrlsrd_modelo)
## Project templating made easy. 

![](recorded.gif)

## Intro
Project templating made easy using Github templates. Think of it as create-react-app for any project type you could think of.

## Usage
1. Create a file called 'config.json' in one of the following paths: [/etc/modelo/ $HOME/.config/modelo .modelo/]
2. Create a personal access token on Github with read access to repositories
3. Add the following content: 
``` json
        { "username": "<github username>", "token": "<github token>" } 
```
* Note: Only 'template repositories' will be listed in the CLI

## TODO
* Allow users to use other Github users templates

## What does Modelo mean?
Modelo is Portuguese for template. I chose the name because I have a good friend from Portugal, who recently lost a loved one due to cancer.
