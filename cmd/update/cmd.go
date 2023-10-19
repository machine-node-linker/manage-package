package update

import (
	"fmt"

	"github.com/spf13/cobra"

	flags "github.com/machine-node-linker/manage-package/pkg/cmd"
	"github.com/machine-node-linker/manage-package/pkg/cmd/update"
	"github.com/machine-node-linker/manage-package/pkg/lib/file"
)

func NewCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update --file filename ",
		Short: "update olm.package schema file",
		Long:  "CLI to update olm.package schema file for operator-framework/operator-registry",
		PreRunE: func(_ *cobra.Command, _ []string) error {
			for _, filename := range flags.UpdateFiles {
				if err := file.CheckFileVar(filename); err != nil {
					return fmt.Errorf("file error: %w", err)
				}
			}
			return nil
		},
		Args: cobra.NoArgs,
		RunE: update.Run,
	}

	cmd.Flags().StringVarP(&flags.IconFile, "icon", "i", "", "icon file to add to package")
	cmd.Flags().StringVarP(&flags.DescriptionFile, "description", "d", "", "description file to add to package")
	cmd.Flags().StringVarP(&flags.PackageFile, "file", "f", "", "package file")

	cmd.MarkFlagRequired("file")
	cmd.MarkFlagFilename("file")
	cmd.MarkFlagFilename("icon")
	cmd.MarkFlagFilename("description")

	return cmd
}
