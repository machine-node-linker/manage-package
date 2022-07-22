package main

import (
	"os"

	"github.com/machine-node-linker/manage-package/cmd/root"
	"github.com/machine-node-linker/manage-package/pkg/github"
)

func main() {
	cmd := root.NewCMD()

	if _, ok := os.LookupEnv("GITHUB_ACTIONS"); ok {
		cmd.SetErr(github.ErrorWriter{})
	}

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
