package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/avevlad/gdotfiles/internal/constants"
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

func (cfg *Config) Sync() {
	content, _ := json.MarshalIndent(cfg, "", " ")
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

	if err := json.Unmarshal(data, &cfg); err != nil {
		log.Fatal().Err(err).Msg("unmarshal cfg file")
	}
}

func (cfg Config) GetReposUrls() []string {
	return []string{
		cfg.GithubIgnoreGitUrl,
		cfg.ToptalIgnoreGitUrl,
		cfg.GitattributeGitUrl,
	}
}

func (cfg Config) GetReposFolders() (folders []string) {
	for _, url := range cfg.GetReposUrls() {
		split := strings.Split(url, "/")
		folder := strings.Join(split[len(split)-2:], "_")
		folders = append(folders, folder)
	}

	return folders
}

func (cfg Config) getReposFoldersWithToptalHotfix() (folders []string) {
	for _, fld := range cfg.GetReposFolders() {
		if strings.Contains(fld, "toptal") {
			fld = path.Join(fld, "templates")
		}
		folders = append(folders, fld)
	}

	return folders
}

func (cfg Config) GetReposFoldersWithCustomFolder() (folders []string) {
	folders = append(folders, constants.CustomFolder)
	folders = append(folders, cfg.getReposFoldersWithToptalHotfix()...)

	return folders
}
