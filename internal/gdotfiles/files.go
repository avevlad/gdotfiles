package gdotfiles

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/avevlad/gdotfiles/internal/config"
	"github.com/avevlad/gdotfiles/internal/utils"
)

type File struct {
	name   string
	folder string
}

type Files struct {
	list       []File
	nameMaxTpl string
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
		if strings.Contains(v, "toptal") {
			v = path.Join(v, "templates")
		}
		files, _ := ioutil.ReadDir(path.Join(utils.UserConfigDir(), v))
		m[v] = files
	}

	for folder, files := range m {
		for _, f := range files {
			if f.IsDir() || f.Name() == ".gitattribute" || f.Name() == ".gitignore" {
				continue
			}
			if !strings.Contains(f.Name(), ".gitignore") && !strings.Contains(f.Name(), ".gitattribute") {
				continue
			}
			len := len(f.Name())
			if len > maxLength {
				maxLength = len
			}
			// fmt.Println(f.Name(), f.Mode().Type())
			fls.list = append(fls.list, File{name: f.Name(), folder: folder})
		}
	}

	for i := 0; i < maxLength; i++ {
		fls.nameMaxTpl += " "
	}

	fmt.Println("files", len(fls.list))
}
