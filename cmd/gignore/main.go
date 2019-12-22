package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/avevlad/gignore/internal/logger"
	"github.com/rs/zerolog/log"
)

func testLog() {
	fmt.Println("run test log")
	log.Debug().Msg("dbg22222")
}

func runFZF(input []string) string {
	bufOut := new(bytes.Buffer)
	cmd := exec.Command("s2h", "-c", "fzf")

	cmd.Stdin = strings.NewReader(strings.Join(input, "\n"))
	cmd.Stdout = bufOut
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal().Err(err).Msg("runFZF")
	}

	fmt.Println(strings.TrimSpace(string(bufOut.Bytes())) == "bar")

	return string(bufOut.Bytes())
}

func main() {
	debug := flag.Bool("debug", false, "sets log level to debug")
	logger.InitLogger()
	log.Debug().Bool("debug", *debug).Send()
	testLog()
	customErr := errors.New("custom fatal error")

	// utils.RemoveDuplicateStr([]string{})
	log.Fatal().Err(customErr).Msg("fatal msg")
}
