// Package fileio provides basic functions for file operations.
package fileio_test

import (
	"bytes"
	"fileio"
	"os"
	"reflect"
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
		{
			name: "empty path",
			args: args{
				filePath: "",
				contents: []byte("abc"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if err := fileio.GenFile(tt.args.filePath, tt.args.contents); err != nil {
				if !tt.wantErr {
					t.Errorf("GenFile() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			defer os.Remove(tt.args.filePath)
			c, err := fileio.FileContents(tt.args.filePath)
			if err != nil {
				t.Errorf("FileContents error = %v", err)
				return
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
				t.Errorf("FileContents error = %v", err)
				return
			}
			if bytes.Compare(c, tt.contents) != 0 {
				t.Errorf("GenFile() contents = %v, want = %v", c, tt.contents)
			}
		})
	}
}

func TestFileContents(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "simple",
			args: args{
				filePath: "TestFileContents",
			},
			want:    []byte("abc"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// test file setup
			if err := fileio.GenFile(tt.args.filePath, tt.want); err != nil {
				t.Fatalf("GenFile error: %v", err)
			}
			defer os.Remove(tt.args.filePath)

			got, err := fileio.FileContents(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileContents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FileContents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBindFiles(t *testing.T) {
	type args struct {
		srcNames []string
		dstName  string
	}
	tests := []struct {
		name        string
		args        args
		srcContents [][]byte
		wantErr     bool
	}{
		{
			name: "simple",
			args: args{
				srcNames: []string{
					"TestBindFilesSrc1", "TestBindFilesSrc2", "TestBindFilesSrc3",
				},
				dstName: "TestBindFilesDst",
			},
			srcContents: [][]byte{
				[]byte("Yabu"), []byte("kara"), []byte("stick"),
			},
			wantErr: false,
		},
		{
			name: "empty dstPath",
			args: args{
				srcNames: []string{
					"TestBindFilesSrc3", "TestBindFilesSrc4",
				},
				dstName: "",
			},
			srcContents: [][]byte{
				[]byte("Isino"), []byte("uenimo"), []byte("3years"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// files setup
			for i, v := range tt.args.srcNames {
				if err := fileio.GenFile(v, tt.srcContents[i]); err != nil {
					t.Fatalf("GenFile fail: %v. leave files: %v", err, tt.args.srcNames[:i])
				}
			}
			defer func() {
				for _, v := range tt.args.srcNames {
					os.Remove(v)
				}
				os.Remove(tt.args.dstName)
			}()

			if err := fileio.BindFiles(tt.args.srcNames, tt.args.dstName); err != nil {
				if !tt.wantErr {
					t.Errorf("BindFiles() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			// contents check
			var wantContents []byte
			for _, v := range tt.srcContents {
				wantContents = append(wantContents, v...)
			}
			gotContents, err := fileio.FileContents(tt.args.dstName)
			if err != nil {
				t.Errorf("FileContents() fail: %v", err)
				return
			}
			if !reflect.DeepEqual(gotContents, wantContents) {
				t.Errorf("gotContents = %v, wantContents = %v", gotContents, wantContents)
				return
			}
		})
	}
}
