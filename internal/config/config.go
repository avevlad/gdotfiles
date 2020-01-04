package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/avevlad/gdotfiles/internal/utils"
	"github.com/rs/zerolog/log"
)

type Config struct {
	GithubIgnoreGitUrl string
	ToptalIgnoreGitUrl string
	GitattributeGitUrl string

	CustomGitFilesFolderPath string
}

func NewConfig() *Config {
	return &Config{
		GithubIgnoreGitUrl: "https://github.com/github/gitignore",
		ToptalIgnoreGitUrl: "https://github.com/toptal/gitignore",
		GitattributeGitUrl: "https://github.com/alexkaratarakis/gitattributes",

		CustomGitFilesFolderPath: utils.GetCustomGitFilesFolderPath(),
	}
}

func (c *Config) Sync() {
	content, _ := json.MarshalIndent(c, "", " ")
	data, err := ioutil.ReadFile(utils.GetConfigPath())
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if err := ioutil.WriteFile(utils.GetConfigPath(), content, 0644); err != nil {
				log.Fatal().Err(err).Msg("write cfg file")
			}
		} else {
			log.Fatal().Err(err).Msg("read cfg file")
		}

		return
	}

	if err := json.Unmarshal(data, &c); err != nil {
		log.Fatal().Err(err).Msg("unmarshal cfg file")
	}
}
