package main

import (
	"fmt"
	"os"
	"github.com/Pandaman331/SysAdminTools/commands"
)

func main() {
	arg := os.Args[1:]

	if len(arg) == 0 {
		fmt.Println("No arguments provided.")
		os.Exit(0)
	} else if arg[0] == "-h" || arg[0] == "--help" {
		commands.ShowHelp()
		os.Exit(0)
	}

	os.Exit(1)
}
