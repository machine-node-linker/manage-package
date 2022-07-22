package update

import (
	"fmt"

	"github.com/machine-node-linker/manage-package/pkg/lib/schema/pkg"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, _ []string) error {
	file, _ := cmd.Flags().GetString("file")
	icon, _ := cmd.Flags().GetString("icon")
	description, _ := cmd.Flags().GetString("description")

	pkgobj, err := pkg.LoadFile(file)
	if err != nil {
		return fmt.Errorf("Unable to load file: %w", err)
	}
	if icon != "" {
		pkgobj.AddIcon(icon)
	}
	if description != "" {
		pkgobj.AddDescription(description)
	}

	if err = pkgobj.WriteToFile(file); err != nil {
		return fmt.Errorf("Unable to write package file: %w", err)
	}

	return nil
}
