package main

import (
	"flag"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

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

	log.Print("Some msg")
	log.Debug().
		Str("SomeStr", "str_value").
		Float64("float", 1337.88).
		Bool("debug", *debug).
		Msg("Hello World")
}
