package gignore

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/avevlad/gignore/internal/utils"
	"github.com/rs/zerolog/log"
)

func Run() error {
	fmt.Println("Run app")
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
