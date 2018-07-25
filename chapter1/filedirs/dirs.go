package filedirs

import (
	"os"
	"github.com/pkg/errors"
	"io"
)

// Operate manipulates files and directories
func Operate() error {
	if err := os.Mkdir("example_dir", os.FileMode(0755)); err != nil {
		return err
	}

	if err := os.Chdir("example_dir"); err != nil {
		return err
	}

	// f is a generic file object, it also implements
	// multiple interfaces and can be used as a reader
	// or writer if the correct bits are set when opening
	f, err := os.Create("test.txt")
	if err != nil {
		return err
	}

	// we write a known-length value to the file and validate
	// that it wrote correctly
	value := []byte("hellon")
	count, err := f.Write(value)
	if err != nil {
		return err
	}
	if count != len(value) {
		return errors.New("incorrect length returned from write")
	}

	if err := f.Close(); err != nil {
		return err
	}

	// read the file
	f, err = os.Open("test.txt")
	if err != nil {
		return err
	}

	io.Copy(os.Stdout, f)
	if err := f.Close(); err != nil {
		return err
	}

	// go to the /tmp directory
	if err := os.Chdir(".."); err != nil {
		return err
	}

	// cleanup: careful can be dangerous!
	if err := os.RemoveAll("example_dir"); err != nil {
		return err
	}

	return nil
}