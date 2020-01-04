package main

import (
	"os"

	"github.com/avevlad/gignore/internal/gignore"
)

func main() {
	err := gignore.Run()
	if err != nil {
		os.Exit(1)
	}
}
