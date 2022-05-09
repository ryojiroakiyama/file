// Package fileio provides basic functions for file operations.
package fileio_test

import (
	"bytes"
	"fileio"
	"os"
	"testing"
)

func TestGenFile(t *testing.T) {
	type args struct {
		filePath string
		contents []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "simple",
			args: args{
				filePath: "simple",
				contents: []byte("abc"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if err := fileio.GenFile(tt.args.filePath, tt.args.contents); (err != nil) != tt.wantErr {
				t.Errorf("GenFile() error = %v, wantErr %v", err, tt.wantErr)
			}
			defer os.Remove(tt.args.filePath)
			c, err := fileio.FileContents(tt.args.filePath)
			if err != nil {
				t.Fatalf("FileContents error = %v", err)
			}
			if bytes.Compare(c, tt.args.contents) != 0 {
				t.Errorf("GenFile() contents = %v, want = %v", c, tt.args.contents)
			}
		})
	}
}

func TestGenTmpFile(t *testing.T) {
	tests := []struct {
		name     string
		contents []byte
		wantErr  bool
	}{
		{
			name:     "simple",
			contents: []byte("abc"),
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			gotfilePath, err := fileio.GenTmpFile(bytes.NewReader(tt.contents))
			if (err != nil) != tt.wantErr {
				t.Errorf("GenTmpFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			defer os.Remove(gotfilePath)
			c, err := fileio.FileContents(gotfilePath)
			if err != nil {
				t.Fatalf("FileContents error = %v", err)
			}
			if bytes.Compare(c, tt.contents) != 0 {
				t.Errorf("GenFile() contents = %v, want = %v", c, tt.contents)
			}
		})
	}
}
