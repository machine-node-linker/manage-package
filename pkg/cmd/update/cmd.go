package update

import (
	"fmt"

	flags "github.com/machine-node-linker/manage-package/pkg/cmd"
	"github.com/machine-node-linker/manage-package/pkg/lib/schema"
	"github.com/machine-node-linker/manage-package/pkg/log"
	"github.com/spf13/cobra"
)

func Run(_ *cobra.Command, _ []string) error {
	pkg, err := schema.LoadPackageFile(&flags.PackageFile)
	if err != nil {
		return fmt.Errorf("Unable to load file: %w", err)
	}

	switch {
	case flags.IconFile != "":
		err = pkg.AddIcon(&flags.IconFile)
		if err != nil {
			return fmt.Errorf("Unable to add icon: %w", err)
		}
	case pkg.Icon != nil:
		log.Debug.Println("Iconfile not specified, using existing Package icon")
	default:
		log.Debug.Println("Iconfile not specified and not set in existing Package")
	}

	switch {
	case flags.DescriptionFile != "":
		err = pkg.AddDescription(&flags.DescriptionFile)
		if err != nil {
			return fmt.Errorf("Unable to add description: %w", err)
		}
	case pkg.Description != "":
		log.Debug.Println("Descripton file not specified, using existing Package description")
	default:
		return fmt.Errorf("Desciption file not specified and not set in existing package")
	}

	if err = pkg.WriteToFile(&flags.PackageFile); err != nil {
		return fmt.Errorf("Unable to write package file: %w", err)
	}

	return nil
}
