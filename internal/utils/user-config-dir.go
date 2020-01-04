package utils

import (
	"os"
	"path/filepath"

	"github.com/avevlad/gdotfiles/internal/constants"
	"github.com/rs/zerolog/log"
)

func UserConfigDir() string {
	dir, err := os.UserConfigDir()

	if err != nil {
		log.Fatal().Err(err).Msg("UserConfigDir is not defined")
	}

	return filepath.Join(dir, constants.AppName)
}

func GetCustomGitFilesFolderPath() string {
	return filepath.Join(UserConfigDir(), "custom-git-files-path")
}

func GetConfigPath() string {
	return filepath.Join(UserConfigDir(), "config.json")
}

func MakeDirIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Debug().Str("path", path).Msg("MkdirAll")
		return os.MkdirAll(path, os.ModePerm)
	}
	return nil
}
