package git

import (
	"reflect"
	"testing"
)

func TestGists_GetFiles(t *testing.T) {
	tests := []struct {
		name  string
		gists Gists
		want  Files
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.gists.GetFiles(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Gists.GetFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}
