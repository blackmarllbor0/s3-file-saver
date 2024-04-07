package logger

import (
	"io"
	"os"
)

type MultiWriter struct {
	File   *os.File
	Stdout io.Writer
}

func (mv *MultiWriter) Write(p []byte) (int, error) {
	if _, err := mv.Stdout.Write(p); err != nil {
		return 0, err
	}

	return mv.File.Write(p)
}
