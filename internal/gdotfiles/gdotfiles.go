package gdotfiles

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/avevlad/gdotfiles/internal/build"
	"github.com/avevlad/gdotfiles/internal/constants"

	"github.com/avevlad/gdotfiles/internal/utils"
	"github.com/rs/zerolog/log"
)

func Run() error {
	parseOsArgs()
	parseFlags()

	fmt.Println("Run app")
	setupDataDirs()
	println("----")
	println(build.Revision)
	println(build.Version)
	println("----")

	fmt.Println("CheckFzfExist", utils.CheckFzfExist())
	fmt.Println("CheckGitExist", utils.CheckGitExist())
	return nil
}

func runFZF(input []string) string {
	bufOut := new(bytes.Buffer)
	cmd := exec.Command("sh", "-c", "fzf")

	cmd.Stdin = strings.NewReader(strings.Join(input, "\n"))
	cmd.Stdout = bufOut
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal().Err(err).Msg("runFZF")
	}

	fmt.Println(strings.TrimSpace(string(bufOut.Bytes())) == "bar")

	return string(bufOut.Bytes())
}

func setupDataDirs() {
	appDir := utils.UserConfigDir()
	if err := utils.MakeDirIfNotExists(appDir); err != nil {
		log.Fatal().Err(err).Msg("setupDataDirs")
	}
}

func parseOsArgs() {
	osArgs := os.Args[1:]

	for _, arg := range osArgs {
		switch {
		case strings.HasPrefix(arg, "--flag="):
			flagVal := arg[len("--flag="):]
			fmt.Println("flagVal", flagVal)
		case arg == "--version", arg == "version", arg == "-v":
			fmt.Printf("%s\n", build.Version+" ("+build.Revision+")")
			os.Exit(0)
		case arg == "-h", arg == "-help", arg == "--help":
			fmt.Print(helpText())
			os.Exit(0)
		default:
		}
	}
}

func parseFlags() {
	opts := struct {
		Name string
		Type string
		From string
	}{}

	flag.StringVar(&opts.Name, "name", "", "")
	flag.StringVar(&opts.Type, "type", "gitignore", "")
	flag.StringVar(&opts.From, "from", "github", "")

	flag.Usage = func() {
		// print(helpText())
	}

	flag.Parse()
	fmt.Println("opts", opts)
	fmt.Println("tail:", flag.Args())
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
