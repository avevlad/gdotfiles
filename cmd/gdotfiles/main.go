package main

import (
	"flag"
	"fmt"
	"github.com/avevlad/gdotfiles/internal/build"
	"github.com/avevlad/gdotfiles/internal/constants"
	"github.com/rs/zerolog/log"
	"os"
	"strings"

	"github.com/avevlad/gdotfiles/internal/gdotfiles"
)

func main() {
	//oldArgs := os.Args
	//defer func() { os.Args = oldArgs }()
	//os.Args = []string{"???", "--name=FooBar", "--from=gh", "--verbose"}
	fmt.Println(os.Args)
	fmt.Println(os.Args[1:])

	buildRestFlags(os.Args[1:])

	app := gdotfiles.NewApp()
	flag.Parse()
	err := app.Run()
	if err != nil {
		log.Fatal().Err(err).Send()
		os.Exit(1)
	}
}

func buildRestFlags(osArgs []string) (hasVerbose bool) {
	for _, arg := range osArgs {
		switch {
		case strings.HasPrefix(arg, "--flag="):
			flagVal := arg[len("--flag="):]
			fmt.Println("flagVal", flagVal)
		case arg == "--verbose":
			hasVerbose = true
		case arg == "--version", arg == "version", arg == "-v":
			fmt.Printf("%s\n", build.Version+" ("+build.Revision+")")
			os.Exit(0)
		case arg == "-h", arg == "-help", arg == "--help":
			fmt.Print(helpText())
			os.Exit(0)
		default:
		}
	}

	return hasVerbose
}

var helpText = func() string {
	return `
Usage:
	` + constants.AppName + ` [options]

Options:
	--name=...        Name of template (Node | Scala | Symfony | 1C-Bitrix...)
	--type=...        Type of git file (ignore | attributes, default ignore)
	--from=...        Source (github | toptal | local | alexkaratarakis, default
	                  github or alexkaratarakis)
	--verbose         Print all logs output
	--yes             Automatic yes to prompts

Examples:
	# Automatic detect project language and choice .gitignore from several options (depends on fzf)
	gdotfiles

	# Create Scala .gitignore file from github.com/github/gitignore
	gdotfiles --name=Scala --from=github
	
	# Create C++ .gitattributes file from github.com/alexkaratarakis/gitattributes
	gdotfiles --name=C++ --type=attributes

	# Create two gitignore templates in one .gitignore file from github.com/github/gitignore
	gdotfiles --name=Scala
`
}
