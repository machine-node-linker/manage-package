package github

import (
	"fmt"
	"io"
)

const ErrorCommand = "::error::"
const DebugCommand = "::debug::"
const GithubLogFlag = 0

type ErrorWriter struct {
	io.Writer
}

func (ErrorWriter) Write(p []byte) (int, error) {
	n, err := fmt.Printf("%s%s", ErrorCommand, p)

	if n < len(ErrorCommand) {
		n = 0
	} else {
		n -= len(ErrorCommand)
	}

	if err != nil {
		err = fmt.Errorf("unable to write to stderr: %w", err)
	}

	return n, err
}
