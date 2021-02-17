package mvn

type PomConfig struct {
	File         string
	Automatic    bool
	KeepSnapshot bool
	Mayor        string
	Minor        string
	Patch        string
}
