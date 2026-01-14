package main

var (
	version   = "dev"
	commit    = "none"
	buildTime = "unknown"
)

func Version() string {
	return version
}

func String() string {
	return version + " " + buildTime + " (" + commit + ")"
}
