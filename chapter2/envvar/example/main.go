package main

import (
	"io/ioutil"
	"os"
	"bytes"
	"cookbook/chapter2/envvar"
	"fmt"
)

// Config will hold the config we capture
// from the json file and env vars
type Config struct {
	Version string `json:"version" required:"true"`
	IsSafe bool `json:"is_safe" default:"true"`
	Secret string `json:"secret"`
}

func main() {
	var err error

	// create a temporary file to hold
	// an example json file
	tmpFile, err := ioutil.TempFile("", "tmp")
	if err != nil {
		panic(err)
	}
	defer tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	// create a json file to hold our secrets
	secrets := `{"secret": "so so secret"}`

	if _, err := tmpFile.Write(bytes.NewBufferString(secrets).Bytes()); err != nil {
		panic(err)
	}

	// we can easily set env variables as needed
	if err = os.Setenv("EXAMPLE_VERSION", "1.0.0"); err != nil {
		panic(err)
	}
	if err = os.Setenv("EXAMPLE_ISSAFE", "false"); err != nil {
		panic(err)
	}

	c := Config{}
	if err = envvar.LoadConfig(tmpFile.Name(), "EXAMPLE", &c); err != nil {
		panic(err)
	}

	fmt.Println("secrets file contains =", secrets)

	// we can also read them
	fmt.Println("EXAMPLE_VERSION =", os.Getenv("EXAMPLE_VERSION"))
	fmt.Println("EXAMPLE_ISSAFE =", os.Getenv("EXAMPLE_ISSAFE"))

	// the final config is a mix of json and env vars
	fmt.Printf("Final Config: %#v\n", c)
}
