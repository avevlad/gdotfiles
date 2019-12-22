package main

import (
	"fmt"
	"os"

	"github.com/avevlad/gignore/internal/gignore"
)

var revision = "unknown"

func main() {
	fmt.Println("revision", revision)
	err := gignore.Run()
	if err != nil {
		os.Exit(1)
	}
}
