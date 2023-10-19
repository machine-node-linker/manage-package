package file

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/machine-node-linker/manage-package/pkg/log"
)

func CheckFileVar(filename *string) error {
	if _, err := os.Stat(*filename); *filename != "" && err != nil {
		log.Debug.Printf("unable to read %s", *filename)
		debugLogPaths()

		return fmt.Errorf("unable to read file: %w", err)
	}

	return nil
}

func debugLogPaths() {
	root := "."
	fileSystem := os.DirFS(root)

	if err := fs.WalkDir(fileSystem, ".", func(path string, _ fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		log.Debug.Println(path)
		return nil
	}); err != nil {
		log.Debug.Printf("Unable to print fs: %v", err)
	}
}
