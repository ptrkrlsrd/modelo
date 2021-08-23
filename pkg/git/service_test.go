package git

import (
	"context"
	"reflect"
	"testing"

	"github.com/shurcooL/githubv4"
)

type mockClient struct{}

func newMockGistQuery() GistQuery {
	gists := Gists{Gist{Name: "testfile", Files: Files{File{Extension: "go", Text: "package main"}}}}
	gistData := gistData{
		Nodes: gists,
	}

	return GistQuery{
		Viewer: gistViewer{
			Gists: gistData,
		},
	}
}

func newMockRepositoryQuery() RepositoryQuery {
	repos := Repositories{Repository{Name: "testfile", URL: "https://github.com/repo/repo.git"}}
	repoData := repositoryData{
		Nodes: repos,
	}

	return RepositoryQuery{
		Viewer: repositoryViewer{
			Repositories: repoData,
		},
	}
}

func (c *mockClient) Query(ctx context.Context, q interface{}, variables map[string]interface{}) error {
	switch t := q.(type) {
	case *GistQuery:
		*t = newMockGistQuery()
	case *RepositoryQuery:
		*t = newMockRepositoryQuery()
	}

	return nil
}

func (c mockClient) Mutate(ctx context.Context, m interface{}, input githubv4.Input, variables map[string]interface{}) error {
	return nil
}

func TestNewService(t *testing.T) {
	type args struct {
		githubUsername string
		githubToken    string
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		{
			name: "Can make a new service",
			args: args{githubUsername: "github_username", githubToken: "github_token"},
			want: Service{
				GithubClient: newClient("github_token"),
				auth: Auth{
					Username: "github_username",
					Token:    "github_token",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.githubUsername, tt.args.githubToken); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetUsersGists(t *testing.T) {
	type fields struct {
		client mockClient
		auth   Auth
	}

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Gists
		wantErr bool
	}{
		{
			name: "can get gists",
			fields: fields{
				client: mockClient{},
				auth:   Auth{},
			},
			want: newMockGistQuery().Viewer.Gists.Nodes,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := Service{
				GithubClient: &tt.fields.client,
				auth:         tt.fields.auth,
			}
			got, err := service.GetUsersGists(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetGists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetGists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetRepositories(t *testing.T) {
	type fields struct {
		client RepoReaderWriter
		auth   Auth
	}
	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Repositories
		wantErr bool
	}{
		{
			name: "can get repos",
			fields: fields{
				client: &mockClient{},
				auth:   Auth{},
			},
			want: newMockRepositoryQuery().Viewer.Repositories.Nodes,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := Service{
				GithubClient: tt.fields.client,
				auth:         tt.fields.auth,
			}
			got, err := service.GetRepositories(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetRepositories() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetRepositories() = %v, want %v", got, tt.want)
			}
		})
	}
}
