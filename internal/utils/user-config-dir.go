package utils

import (
	"os"
	"path/filepath"

	"github.com/avevlad/gignore/internal/constants"
	"github.com/rs/zerolog/log"
)

func UserConfigDir() string {
	dir, err := os.UserConfigDir()

	if err != nil {
		log.Fatal().Err(err).Msg("UserConfigDir is not defined")
	}

	return filepath.Join(dir, constants.AppName)
}

func MakeDirIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, os.ModePerm)
	}
	return nil
}
