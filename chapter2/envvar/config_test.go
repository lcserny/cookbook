package envvar

import (
	"testing"
	"io/ioutil"
	"os"
	"bytes"
)

func TestLoadConfig(t *testing.T) {
	c := struct {
		Example string `json:"example"`
	}{}

	type args struct {
		path string
		envPrefix string
		config interface{}
	}

	tests := []struct{
		name string
		args args
		wantErr bool
	}{
		{"base-case", args{"123", "", &c}, true},
		{"no file", args{"", "", &c}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LoadConfig(tt.args.path, tt.args.envPrefix, tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("LoadConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoadFile(t *testing.T) {
	c := struct {
		Example string `json:"example"`
	}{}

	tmpFile, err := ioutil.TempFile("", "tmp")
	if err != nil {
		panic(err)
	}
	defer tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write(bytes.NewBufferString("{}").Bytes()); err != nil {
		t.Errorf("Error writing into tmp file, err: %v", err)
	}

	type args struct {
		path string
		config interface{}
	}

	tests := []struct{
		name string
		args args
		wantErr bool
	}{
		{"base-case", args{"", &c}, true},
		{"base-case", args{tmpFile.Name(), &c}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LoadFile(tt.args.path, tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("LoadFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
