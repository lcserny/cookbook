package tempfiles

import (
	"io/ioutil"
	"os"
	"fmt"
)

// WorkWithTemp will give some basic patterns for working
// with temporary files and directories
func WorkWithTemp() error {
	// if you need a temporary place to store files with the same name,
	// a temp directory is a good way to approach it, the first argument
	// being blank means it will create the directory in the location
	// defined by os.TempDir()
	t, err := ioutil.TempDir("", "tmp")
	if err != nil {
		return err
	}

	// this will delete everything inside the temp file
	// when this function exists if you want to do this later,
	// be sure to return the dir name to the calling function
	defer os.RemoveAll(t)

	// the directory must exist to create the temp file
	// t is an *os.File object
	tf, err := ioutil.TempFile(t, "tmp")
	if err != nil {
		return err
	}

	fmt.Println(tf.Name())

	// normally we would delete the temp file here
	// but because we are placing it in a temp dir
	// it gets cleaned up by the earlier defer

	return nil
}
