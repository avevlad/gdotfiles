package utils

import (
	"os/exec"
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
