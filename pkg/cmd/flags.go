package cmd

var (
	Verbose         bool
	PackageFile     string
	IconFile        string
	DescriptionFile string
	UpdateFiles     = []*string{&PackageFile, &IconFile, &DescriptionFile}
)
