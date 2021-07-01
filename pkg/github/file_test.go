package github

import (
	"reflect"
	"testing"
)

func TestFiles_ToMap(t *testing.T) {
	testFile := File{
		Name:      "test",
		Extension: "go",
		Text:      "package main",
	}

	wantFileMap := make(map[string]File)
	wantFileMap[testFile.Name] = testFile

	files := Files{}
	files = append(files, testFile)

	tests := []struct {
		name    string
		files   Files
		want    map[string]File
		wantErr bool
	}{
		{
			name:    "Valid file",
			files:   files,
			wantErr: false,
			want: map[string]File{"test": {
				Name:      "test",
				Extension: "go",
				Text:      "package main",
			},
			},
		},
		{
			name:    "Invalid name",
			wantErr: true,
			files: Files{
				{
					Name: "",
				},
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.files.ToMap()
			if (err != nil) != tt.wantErr {
				t.Errorf("Files.ToMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Files.ToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFiles_GetNames(t *testing.T) {
	tests := []struct {
		name  string
		files Files
		want  []string
	}{
		{
			name: "Can get names",
			files: []File{
				{
					Name: "name",
				},
			},
			want: []string{"name"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.files.GetNames(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Files.GetNames() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFiles_Filter(t *testing.T) {
	type args struct {
		ignored []string
	}

	tests := []struct {
		name  string
		files Files
		args  args
		want  Files
	}{
		{
			name: "Filters names",
			files: []File{
				{
					Name: "Ok",
				},
				{
					Name: "Not Ok",
				},
			},
			args: args{
				ignored: []string{
					"Not Ok",
				},
			},
			want: []File{
				{
					Name: "Ok",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.files.Filter(tt.args.ignored); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Files.Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_Write(t *testing.T) {
	type fields struct {
		Name      string
		Extension string
		Text      string
	}
	type args struct {
		filePath string
		fileName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := &File{
				Name:      tt.fields.Name,
				Extension: tt.fields.Extension,
				Text:      tt.fields.Text,
			}
			if err := file.Write(tt.args.filePath, tt.args.fileName); (err != nil) != tt.wantErr {
				t.Errorf("File.Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
