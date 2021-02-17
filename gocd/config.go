package gocd

type BuildConfig struct {
	Project_Dir string `yaml: "Project_Dir"`
	Settings_File string `yaml: Settings_File`
	Java_Version string `yaml:"Java_Version"`

}
