package logger

import (
	"os"
	"time"

	"github.com/avevlad/gdotfiles/internal/constants"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type ConsoleLoggerOpts struct {
	Level zerolog.Level
}

type CustomWriter struct {
	lw    zerolog.LevelWriter
	level zerolog.Level
}

func (w *CustomWriter) Write(p []byte) (n int, err error) {
	return w.lw.Write(p)
}

func (w *CustomWriter) WriteLevel(level zerolog.Level, p []byte) (n int, err error) {
	if level >= w.level {
		return w.lw.WriteLevel(level, p)
	}
	return len(p), nil
}

// TODO: config
func InitLogger(opts *ConsoleLoggerOpts) {
	filename := "./" + constants.AppName + ".log"
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal().Err(err).Msg("open log file err")
	}
	consoleWriter := zerolog.MultiLevelWriter(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.Kitchen,
	})
	customWriter := &CustomWriter{consoleWriter, opts.Level}
	mw := zerolog.MultiLevelWriter(
		zerolog.ConsoleWriter{
			Out:        file,
			NoColor:    true,
			TimeFormat: time.RFC3339,
		},
		customWriter,
	)
	log.Logger = zerolog.New(mw).With().Timestamp().Logger()
}

func Cleanup() {
	// ...
}
