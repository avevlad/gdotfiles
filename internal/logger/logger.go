package logger

import (
	"github.com/avevlad/gignore/internal/constants"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

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
func InitLogger() {
	filename := "./" + constants.AppName + ".log"
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal().Err(err).Msg("open log file err")
	}
	consoleWriter := zerolog.MultiLevelWriter(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.Kitchen,
	})
	customWriter := &CustomWriter{consoleWriter, zerolog.FatalLevel}
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
