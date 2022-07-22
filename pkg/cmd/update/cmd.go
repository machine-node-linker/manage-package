package update

import (
	"fmt"

	"github.com/machine-node-linker/manage-package/pkg/lib/schema"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, _ []string) error {
	file, _ := cmd.Flags().GetString("file")
	icon, _ := cmd.Flags().GetString("icon")
	description, _ := cmd.Flags().GetString("description")

	pkg, err := schema.LoadPackageFile(file)
	if err != nil {
		return fmt.Errorf("Unable to load file: %w", err)
	}
	if icon != "" {
		pkg.AddIcon(icon)
	}
	if description != "" {
		pkg.AddDescription(description)
	}

	if err = pkg.WriteToFile(file); err != nil {
		return fmt.Errorf("Unable to write package file: %w", err)
	}

	return nil
}
