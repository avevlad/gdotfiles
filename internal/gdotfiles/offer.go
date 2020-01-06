package gdotfiles

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/avevlad/gdotfiles/internal/utils"
	"github.com/rs/zerolog/log"
)

func offerFoundFile(file File, flags *AppFlags) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal().Err(err)
	}
	fileName := "." + file.GetFileType()

	currentGitFile := path.Join(wd, fileName)
	if _, err := os.Stat(currentGitFile); !os.IsNotExist(err) {
		data, _ := ioutil.ReadFile(currentGitFile)
		if strings.TrimSpace(string(data)) == "" {
			writeGitFile(file, false)
			return
		}
		fmt.Println("We found an already existing", currentGitFile, "file")
		fmt.Println("You can append data to an existing file or replace")
		prompt := true
		if !flags.Yes {
			prompt = utils.YesOrNoPrompt("Do you want to append selected file to the current "+fileName, true)
		}

		writeGitFile(file, prompt)
		return
	}
	writeGitFile(file, false)
}

func writeGitFile(file File, isAppend bool) {
	var (
		writeContent []byte
		fileName     = "." + file.GetFileType()
	)
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal().Err(err)
	}
	currentGitFilePath := path.Join(wd, fileName)
	filePath := path.Join(utils.UserConfigDir(), file.Folder, file.Name)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal().Err(err)
	}
	fmt.Println("isAppend", isAppend)
	writeContent = data
	if isAppend {
		currentFile, err := ioutil.ReadFile(currentGitFilePath)
		if err != nil {
			log.Fatal().Err(err)
		}
		writeContent = append(currentFile, []byte("\n\n")...)
		writeContent = append(writeContent, data...)
	}
	// fmt.Println("-----------")
	// fmt.Println("data")
	// fmt.Println(string(writeContent))
	if err := ioutil.WriteFile(currentGitFilePath, writeContent, 0644); err != nil {
		log.Fatal().Err(err).Msg("write git file err")
	}
	// fmt.Println(wd)
}
