package git

import (
	"reflect"
	"testing"
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
