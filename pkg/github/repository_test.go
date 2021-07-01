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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
