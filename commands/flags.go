package commands

import (
	"fmt"
	"flag"
	"os"
	"strings"
)

type Config struct {
	VersionFlag	*bool
	helpFlag	*bool
}

type flagInfo struct {
	longName	string
	shortName	string
	defaultVal	any
	description	string
	flagType	string
}

func ParseFlags() Config {
	config := Config{}
	registeredFlags = nil

	config.VersionFlag  = config.boolFlag("version", "v", false, "Display version of SAT")
	config.helpFlag = config.boolFlag("help", "h", false, "Display help information")

	flag.Usage = showCustomHelp
	flag.Parse()

	if *config.helpFlag {
		showCustomHelp()
	}

	return config
}

var registeredFlags []flagInfo

func (c *Config) stringFlag(long, short, def, desc string) *string {
	registeredFlags = append(registeredFlags, flagInfo{long, short, def, desc, "string"})
	val := flag.String(long, def, "")
	flag.StringVar(val, short, def, "")
	return val
}

func (c *Config) boolFlag(long, short string, def bool, desc string) *bool {
	registeredFlags = append(registeredFlags, flagInfo{long, short, def, desc, ""})
	val := flag.Bool(long, def, "")
	flag.BoolVar(val, short, def, "")
	return val
}

func (c *Config) intFlag(long, short string, def int, desc string) *int {
	registeredFlags = append(registeredFlags, flagInfo{long, short, def, desc, "int"})
	val := flag.Int(long, def, "")
	flag.IntVar(val, short, def, "")
	return val
}

func showCustomHelp() {
	fmt.Println("Usage: sat [options]")
	fmt.Println("Options:")

	maxWidth := 0
	for _, f := range registeredFlags {
		width := len("-" + f.shortName + ", --" + f.longName + " " + f.flagType)
		if width > maxWidth {
			maxWidth = width
		}
	}

	for _, f := range registeredFlags {
		flag := fmt.Sprintf("-%s, --%s", f.shortName, f.longName)
		flagLen := len(flag)
		if f.flagType != "" {
			flagLen = len(flag + " " + f.flagType)
			flag += " " + f.flagType
		}
		padding := strings.Repeat(" ", maxWidth-flagLen+2)
		fmt.Println(" " + flag + padding + f.description)
	}
	os.Exit(0)
}
