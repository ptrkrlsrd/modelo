package github

import (
	"context"
	"reflect"
	"testing"

	"github.com/shurcooL/githubv4"
)

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
			args: args{githubUsername: "abc", githubToken: ""},
			want: Service{},
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

func TestService_GetGists(t *testing.T) {
	type fields struct {
		GithubClient *githubv4.Client
		auth         Auth
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := Service{
				GithubClient: tt.fields.GithubClient,
				auth:         tt.fields.auth,
			}
			got, err := service.GetGists(tt.args.ctx)
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
