package gdotfiles

import (
	"fmt"
	"testing"

	"github.com/avevlad/gdotfiles/internal/config"
	"github.com/avevlad/gdotfiles/internal/test/assert"
)

func TestListOfFiles(t *testing.T) {
	var files Files

	files.Read(*config.NewConfig())
	file := files.FilterByFlags(&AppFlags{Name: "C++"})
	fmt.Println("find", file)
	assert.Equal(t, file.Name, "C++.gitignore")
	assert.Equal(t, file.Folder, "github_gitignore")

	file = files.FilterByFlags(&AppFlags{Name: "C++", From: "toptal"})
	fmt.Println("find", file)
	assert.Equal(t, file.Name, "C++.gitignore")
	assert.Equal(t, file.Folder, "toptal_gitignore/templates")

	file = files.FilterByFlags(&AppFlags{Name: "C++", Type: "gitattributes"})
	fmt.Println("find", file)
	assert.Equal(t, file.Name, "C++.gitattributes")
	assert.Equal(t, file.Folder, "alexkaratarakis_gitattributes")
}
