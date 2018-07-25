package interfaces

import (
	"io"
	"os"
	"fmt"
)

// Copy copies data from in to out directly first,
// then using a buffer. It also writes to stdout
func Copy(in io.ReadSeeker, out io.Writer) error {
	// we write to out but also Stdout
	w := io.MultiWriter(out, os.Stdout)

	// a standard copy, this can be dangerous if
	// there is a lot of data in it
	if _, err := io.Copy(w, in); err != nil {
		return err
	}

	in.Seek(0, 0)

	// buffered write using 64 byte chunks
	buf := make([]byte, 64)
	if _, err := io.CopyBuffer(w, in, buf); err != nil {
		return err
	}

	// print a new line
	fmt.Println()

	return nil
}
