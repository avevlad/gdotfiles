package main

import (
	"os"

	"github.com/avevlad/gdotfiles/internal/gdotfiles"
)

func main() {
	err := gdotfiles.Run()
	if err != nil {
		os.Exit(1)
	}
}
