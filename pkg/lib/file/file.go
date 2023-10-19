package file

import (
	"fmt"
	"os"
)

func CheckFileVar(filename *string) error {
	if _, err := os.Stat(*filename); *filename != "" && err != nil {
		return fmt.Errorf("unable to read file: %w", err)
	}

	return nil
}
