package root

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	updatecmd "github.com/machine-node-linker/manage-package/cmd/update"
	flags "github.com/machine-node-linker/manage-package/pkg/cmd"
	"github.com/machine-node-linker/manage-package/pkg/log"
)

func NewCMD() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "manage-package",
		Version: os.Getenv("VERSION"),
		Short:   "olm.package file update tool",
		Long:    "CLI to update olm.package schema files for operator-framework/operator-registry",
		Args:    cobra.NoArgs,
		CompletionOptions: cobra.CompletionOptions{
			HiddenDefaultCmd: true,
		},
		PersistentPreRun: func(cmd *cobra.Command, _ []string) {
			if flags.Verbose {
				log.Debug.SetOutput(os.Stdout)
			}
			log.Debug.Printf("%s version %s", cmd.Root().Name(), cmd.Root().Version)
			cmd.Flags().Visit(func(f *pflag.Flag) {
				log.Debug.Printf("set flag %s to %s", f.Name, f.Value)
			})
		},
	}

	rootCmd.PersistentFlags().BoolVarP(&flags.Verbose, "verbose", "v", false, "verbose output")

	rootCmd.AddCommand(
		updatecmd.NewCMD(),
	)

	return rootCmd
}
