package schema

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/h2non/filetype"
	"github.com/machine-node-linker/manage-package/pkg/log"

	"github.com/operator-framework/operator-registry/alpha/declcfg"
)

type Package declcfg.Package

func LoadPackageFile(filename *string) (*Package, error) {
	p := &Package{}

	log.Debug.Printf("Reading Package data from %s", *filename)

	f, err := os.ReadFile(*filename)
	if err != nil {
		return nil, fmt.Errorf("read package: %v", err)
	}

	if err := json.Unmarshal(f, &p); err != nil {
		return nil, fmt.Errorf("parse package: %v: %s", err, f)
	}

	log.Debug.Printf("Parsed Package from %s with name: %s", *filename, p.Name)

	return p, nil
}

func (p *Package) AddIcon(filename *string) error {
	iconData, err := os.ReadFile(*filename)
	if err != nil {
		return fmt.Errorf("read icon: %v", err)
	}

	if len(iconData) == 0 {
		log.Debug.Printf("No data found in Iconfile %s", *filename)
		return nil
	}

	iconType, err := filetype.Match(iconData)
	if err != nil {
		return fmt.Errorf("detect icon mediatype: %v", err)
	}

	if iconType.MIME.Type != "image" {
		return fmt.Errorf("detected invalid type %q: not an image", iconType.MIME.Value)
	}

	log.Debug.Printf("Found image of type %s in %s", iconType.MIME.Type, *filename)

	p.Icon = &declcfg.Icon{
		Data:      iconData,
		MediaType: iconType.MIME.Value,
	}

	return nil
}

func (p *Package) AddDescription(filename *string) error {
	descriptionData, err := os.ReadFile(*filename)
	if err != nil {
		return fmt.Errorf("read description: %v", err)
	}

	if len(descriptionData) == 0 {
		log.Debug.Printf("No data found in Description File %s", *filename)
		return nil
	}

	p.Description = string(descriptionData)

	return nil
}

func (p *Package) WriteToFile(filename *string) error {
	data, err := json.Marshal(p)
	if err != nil {
		return fmt.Errorf("encode package: %v", err)
	}

	log.Debug.Printf("Writing Package data to %s", *filename)

	if err := os.WriteFile(*filename, data, 0664); err != nil {
		return fmt.Errorf("write package: %v", err)
	}

	return nil
}
