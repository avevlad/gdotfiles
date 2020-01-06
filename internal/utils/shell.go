package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func CheckGitExist() bool {
	return CheckBinExist("git", "--version")
}

func CheckFzfExist() bool {
	return CheckBinExist("fzf", "--version")
}

func CheckBinExist(name string, arg ...string) bool {
	_, err := exec.Command(name, arg...).CombinedOutput()

	return err == nil
}

func ExecCommand(cmd string) []byte {
	out, _ := exec.Command("bash", "-c", cmd).CombinedOutput()

	return out
}

func YesOrNoPrompt(question string, defaultValue bool) bool {
	ynWord := "Y/n"
	if !defaultValue {
		ynWord = "y/N"
	}

	reader := bufio.NewReader(os.Stdin)
	var resp string

	for {
		fmt.Fprintf(os.Stderr, "%s (%s) ", question, ynWord)
		resp, _ = reader.ReadString('\n')
		resp = strings.ToLower(strings.TrimSpace(resp))
		if resp == "" {
			return defaultValue
		}
		if resp == "y" || resp == "yes" {
			return true
		} else if resp == "n" || resp == "no" {
			return false
		}
	}
}
