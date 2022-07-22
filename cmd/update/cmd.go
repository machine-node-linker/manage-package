package update

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/machine-node-linker/manage-package/pkg/cmd/update"
)

func NewCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update --file filename ",
		Short: "update olm.package schema file",
		Long:  "CLI to update olm.package schema file for operator-framework/operator-registry",
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			file, _ := cmd.Flags().GetString("file")

			if _, err := os.Stat(file); err != nil {
				return fmt.Errorf("unable to read file: %w", err)
			}
			icon, _ := cmd.Flags().GetString("icon")

			if _, err := os.Stat(icon); icon != "" && err != nil {
				return fmt.Errorf("unable to read file: %w", err)
			}
			description, _ := cmd.Flags().GetString("description")

			if _, err := os.Stat(description); description != "" && err != nil {
				return fmt.Errorf("unable to read file: %w", err)
			}

			if icon == "" && description == "" {
				return fmt.Errorf("one of --icon or --description must be preset")
			}
			return nil
		},
		Args: cobra.NoArgs,
		RunE: update.Run,
	}

	cmd.Flags().StringP("icon", "i", "", "icon file to add to package")
	cmd.Flags().StringP("description", "d", "", "description file to add to package")
	cmd.Flags().StringP("file", "f", "", "package file")

	cmd.MarkFlagRequired("file")
	cmd.MarkFlagFilename("file")
	cmd.MarkFlagFilename("icon")
	cmd.MarkFlagFilename("description")

	return cmd
}
