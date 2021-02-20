package library

const (
	Bundler  = "bundler"
	Cargo    = "cargo"
	Composer = "composer"
	Egg      = "egg"
	Npm      = "npm"
	NuGet    = "nuget"
	Pipenv   = "pipenv"
	Poetry   = "poetry"
	Yarn     = "yarn"
	Jar      = "jar"
	Wheel    = "wheel"
)

var (
	IgnoreDirs = []string{"node_modules", "vendor"}
)
