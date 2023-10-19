package main

import (
	"os"

	"github.com/machine-node-linker/manage-package/cmd/root"
	"github.com/machine-node-linker/manage-package/pkg/github"
	"github.com/machine-node-linker/manage-package/pkg/log"
)

func main() {
	cmd := root.NewCMD()

	if _, ok := os.LookupEnv("GITHUB_ACTIONS"); ok {
		cmd.SetErr(github.ErrorWriter{})
		log.Debug.SetPrefix(github.DebugCommand)
		log.Debug.SetFlags(github.GithubLogFlag)
		log.Info.SetFlags(github.GithubLogFlag)
	}

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
