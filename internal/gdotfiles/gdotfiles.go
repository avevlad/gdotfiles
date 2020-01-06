package gdotfiles

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"sync"

	"github.com/avevlad/gdotfiles/internal/build"
	"github.com/avevlad/gdotfiles/internal/config"
	"github.com/avevlad/gdotfiles/internal/constants"
	"github.com/avevlad/gdotfiles/internal/logger"

	"github.com/avevlad/gdotfiles/internal/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type AppFlags struct {
	Name    string
	Type    string
	From    string
	Verbose bool
}

func (af *AppFlags) registerFlags(fs *flag.FlagSet) {
	fs.StringVar(&af.Name, "name", "", "")
	fs.StringVar(&af.Type, "type", "", "")
	fs.StringVar(&af.From, "from", "", "")
	fs.BoolVar(&af.Verbose, "verbose", false, "")

	fs.Usage = func() {
		// print(helpText())
	}
}

func Run() error {
	var (
		cfg      = config.NewConfig()
		files    Files
		appFlags AppFlags
	)

	setupDataDirs()
	cfg.Sync()

	verbose := buildRestFlags()
	var logLevel = zerolog.FatalLevel

	if verbose {
		logLevel = zerolog.DebugLevel
	}
	logger.InitLogger(&logger.ConsoleLoggerOpts{Level: logLevel})

	log.Debug().Msg("some msg")
	log.Info().Strs("version", []string{build.Version, build.Revision}).Send()

	appFlags.registerFlags(flag.CommandLine)
	flag.Parse()

	// println(build.Revision)
	// println(build.Version)

	// fmt.Println("CheckFzfExist", utils.CheckFzfExist())
	// fmt.Println("CheckGitExist", utils.CheckGitExist())

	downloadRepos(*cfg)
	files.Read(*cfg)

	if appFlags.Name != "" {
		r := files.FilterByFlags(&appFlags)
		if r == (File{}) {
			fmt.Println("No files found, try simplifying the arguments")
			os.Exit(0)
		}
		offerFoundFile(r)
		os.Exit(0)
	}

	var interactiveInput []string

	for _, v := range files.List {
		l := len(v.Name)
		left := v.Name + files.NameMaxTpl[l-1:]

		interactiveInput = append(interactiveInput, left+" ["+v.Folder+"]")
	}

	runFZF(interactiveInput)

	fmt.Println("FINISH")
	return nil
}

func runFZF(input []string) string {
	var (
		bufOut = new(bytes.Buffer)
		cmd    = exec.Command("sh", "-c", "fzf")
	)

	cmd.Stdin = strings.NewReader(strings.Join(input, "\n"))
	cmd.Stdout = bufOut
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal().Err(err).Msg("runFZF")
	}

	fmt.Println(strings.TrimSpace(bufOut.String()) == "bar")

	return bufOut.String()
}

func setupDataDirs() {
	appDir := utils.UserConfigDir()
	if err := utils.MakeDirIfNotExists(appDir); err != nil {
		log.Fatal().Err(err).Msg("setupDataDirs")
	}
	if err := utils.MakeDirIfNotExists(utils.GetCustomGitFilesFolderPath()); err != nil {
		log.Fatal().Err(err).Msg("setupDataDirs custom folder")
	}
}

func downloadRepos(cfg config.Config) {
	if _, err := os.Stat(path.Join(utils.UserConfigDir(), "github_gitignore")); !os.IsNotExist(err) {
		// fmt.Println("not the first run")
		return
	}

	errChan := make(chan error)
	wg := sync.WaitGroup{}
	reposList := cfg.GetReposUrls()
	reposFolders := cfg.GetReposFolders()
	fmt.Println("This is the first run, we need some time to clone and cache gitignore and gitattribute files")

	for i, v := range reposList {
		wg.Add(1)
		go func(index int, url string) {
			defer wg.Done()

			folder := reposFolders[index]
			fmt.Println("Start cloning", url, "in", folder)

			log.Debug().Str("folder", folder).Msg("download start")
			cmd := exec.Command(`git`, `clone`, url, folder)
			cmd.Dir = utils.UserConfigDir()
			resp, err := cmd.CombinedOutput()
			if err != nil {
				if len(resp) > 0 {
					fmt.Println("err resp:", string(resp))
				}
				errChan <- err
			}
			log.Debug().Str("folder", folder).Str("resp", string(resp)).Msg("download finish")
		}(i, v)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	err := <-errChan
	if err != nil {
		log.Fatal().Err(err).Msg("git clone fatal err")
	}
}

func buildRestFlags() (hasVerbose bool) {
	osArgs := os.Args[1:]

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
