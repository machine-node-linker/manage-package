package root

import (
	"github.com/spf13/cobra"

	updatecmd "github.com/machine-node-linker/manage-package/cmd/update"
)

func NewCMD() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "manage-package",
		Short: "olm.package file update tool",
		Long:  "CLI to update olm.package schema files for operator-framework/operator-registry",
		Args:  cobra.NoArgs,
		CompletionOptions: cobra.CompletionOptions{
			HiddenDefaultCmd: true,
		},
	}

	rootCmd.AddCommand(
		updatecmd.NewCMD(),
	)

	return rootCmd
}
