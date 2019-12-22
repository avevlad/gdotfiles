package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func runFZF(input []string) string {
	bufOut := new(bytes.Buffer)
	cmd := exec.Command("sh", "-c", "fzf")

	cmd.Stdin = strings.NewReader(strings.Join(input, "\n"))
	cmd.Stdout = bufOut
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal().Err(err)
	}

	fmt.Println(strings.TrimSpace(string(bufOut.Bytes())) == "bar")

	return string(bufOut.Bytes())
}

func main() {
	debug := flag.Bool("debug", false, "sets log level to debug")
	flag := os.O_CREATE | os.O_APPEND | os.O_WRONLY
	file, err := os.OpenFile("./"+AppName+".log", flag, 0644)
	if err != nil {
		// TODO: ...
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}
	defer func() {
		if file != nil {
			_ = file.Close()
		}
	}()
	if err == nil {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: file, NoColor: true, TimeFormat: time.RFC3339})
	}

	log.Debug().
		Bool("debug", *debug).
		Msg("Hello World")

	runFZF([]string{"foo", "bar", "fzf"})
}
