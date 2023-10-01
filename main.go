package main

import (
	"kwx/cmd"
)

var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)

func init() {
	cmd.Version = Version
	cmd.Commit = Commit
	cmd.Date = Date
}

func main() {
	cmd.Execute()
}
