package main

import (
	"os"
	"github.com/Pandaman331/SysAdminTools/commands"
)

func main() {
	config := commands.ParseFlags()

	if *config.VersionFlag {
		commands.ShowVersion()
		os.Exit(0)
	}
	os.Exit(1)
}
