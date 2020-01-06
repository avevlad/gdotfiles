package gdotfiles

import (
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/avevlad/gdotfiles/internal/config"
	"github.com/avevlad/gdotfiles/internal/utils"
	"github.com/rs/zerolog/log"
)

type File struct {
	Name   string
	Folder string
}

type Files struct {
	List       []File
	NameMaxTpl string
}

func NewFiles() *Files {
	return &Files{}
}

func (fls *Files) Read(cfg config.Config) {
	var (
		reposFolders = cfg.GetReposFoldersWithCustomFolder()
		m            = make(map[string][]os.FileInfo, len(reposFolders))
		maxLength    = 0
	)

	for _, v := range reposFolders {
		files, _ := ioutil.ReadDir(path.Join(utils.UserConfigDir(), v))
		m[v] = files
	}

	for _, folder := range reposFolders {
		files := m[folder]
		for _, f := range files {
			if f.IsDir() || f.Name() == ".gitattributes" || f.Name() == ".gitignore" {
				continue
			}
			if !strings.Contains(f.Name(), ".gitignore") && !strings.Contains(f.Name(), ".gitattributes") {
				continue
			}
			len := len(f.Name())
			if len > maxLength {
				maxLength = len
			}
			// fmt.Println(f.Name(), f.Mode().Type())
			fls.List = append(fls.List, File{Name: f.Name(), Folder: folder})
		}
	}

	for i := 0; i < maxLength; i++ {
		fls.NameMaxTpl += " "
	}

	//content, _ := json.Marshal(fls.List)
	//fmt.Println(string(content))
	// fmt.Println(fls.List[0])
	// fmt.Println("files", len(fls.List))
}

func (fls *Files) FilterByFlags(flags *AppFlags) (result File) {
	for _, v := range fls.List {
		if flags.Type != "" &&
			strings.Contains(v.Name, flags.Name) &&
			strings.Contains(v.Name, flags.Type) {
			log.Debug().Msg("FilterByFlags first statement (by type)")
			result = v
			return result
		}
		if flags.From != "" &&
			strings.Contains(v.Name, flags.Name) &&
			strings.Contains(v.Folder, flags.From) {
			log.Debug().Msg("FilterByFlags second statement (by from)")
			result = v
			return result
		}
		if flags.From == "" && flags.Type == "" && strings.Contains(v.Name, flags.Name) {
			log.Debug().Msg("FilterByFlags third statement (default)")
			result = v
			return result
		}
	}

	return result
}
