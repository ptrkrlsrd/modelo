package github

import (
	"reflect"
	"testing"
)

func TestRepositories_GetTemplates(t *testing.T) {
	tests := []struct {
		name         string
		repositories Repositories
		want         Repositories
	}{
		{
			name: "filters out non-template repo",
			repositories: Repositories{
				Repository{
					Name:       "non-template repo",
					IsTemplate: false,
				},
				Repository{
					Name:       "template repo",
					IsTemplate: true,
				},
			},
			want: Repositories{
				Repository{
					Name:       "template repo",
					IsTemplate: true,
				},
			},
		},
		{
			name: "returns empty array if there are no template repos",
			repositories: Repositories{
				Repository{
					Name:       "non-template repo",
					IsTemplate: false,
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.repositories.GetTemplates(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repositories.GetTemplates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepositories_GetNames(t *testing.T) {
	tests := []struct {
		name                string
		repositories        Repositories
		wantRepositoryNames []string
	}{
		{
			name: "get names from repositories name",
			repositories: Repositories{
				{
					Name: "Repo 1",
				},
				{
					Name: "Repo 2",
				},
			},
			wantRepositoryNames: []string{"Repo 1", "Repo 2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRepositoryNames := tt.repositories.GetNames(); !reflect.DeepEqual(gotRepositoryNames, tt.wantRepositoryNames) {
				t.Errorf("Repositories.GetNames() = %v, want %v", gotRepositoryNames, tt.wantRepositoryNames)
			}
		})
	}
}

func TestRepositories_FindByName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name         string
		repositories Repositories
		args         args
		want         Repository
		wantErr      bool
	}{
		{
			name: "can find repo by name",
			repositories: Repositories{
				{
					Name: "Repo 1",
				},
			},
			args: args{
				name: "Repo 1",
			},
			want:    Repository{Name: "Repo 1"},
			wantErr: false,
		},
		{
			name:         "fails when cant find",
			repositories: Repositories{},
			args: args{
				name: "Repo 1",
			},
			want:    Repository{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.repositories.FindByName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repositories.FindByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repositories.FindByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepositories_Filter(t *testing.T) {
	type args struct {
		ignored []string
	}
	tests := []struct {
		name         string
		repositories Repositories
		args         args
		want         Repositories
	}{
		{
			name: "can filter with one ignored repo",
			repositories: Repositories{
				{
					Name: "repo",
				},
				{
					Name: "ignored",
				},
			},
			args: args{
				ignored: []string{"ignored"},
			},
			want: Repositories{
				{
					Name: "repo",
				},
			},
		},
		{
			name: "can filter multiple ignored repos",
			repositories: Repositories{
				{
					Name: "repo",
				},
				{
					Name: "ignored",
				},
				{
					Name: "another ignored repo",
				},
			},
			args: args{
				ignored: []string{"ignored", "another ignored repo"},
			},
			want: Repositories{
				{
					Name: "repo",
				},
			},
		},
		{
			name: "works when there are no ignored repos",
			repositories: Repositories{
				{
					Name: "repo",
				},
				{
					Name: "repo 2",
				},
				{
					Name: "repo 3",
				},
			},
			args: args{
				ignored: []string{},
			},
			want: Repositories{
				{
					Name: "repo",
				},
				{
					Name: "repo 2",
				},
				{
					Name: "repo 3",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.repositories.Filter(tt.args.ignored); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repositories.Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
