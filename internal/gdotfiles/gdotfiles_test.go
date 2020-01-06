package gdotfiles

import (
	"flag"
	"github.com/avevlad/gdotfiles/internal/test/assert"
	"github.com/avevlad/gdotfiles/internal/utils"
	"os"
	"strings"
	"testing"
)

func readGitignore() (result string) {
	return string(utils.ExecCommand(`cat .gitignore`))
}
func readGitAttributes() (result string) {
	return string(utils.ExecCommand(`cat .gitattributes`))
}

func rmGitFiles() {
	utils.ExecCommand(`rm .gitattributes`)
	utils.ExecCommand(`rm .gitignore`)
}

func TestGithubNode(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	rmGitFiles()
	os.Args = []string{"???", "--name=Node", "--yes", "--verbose"}
	app := NewApp()
	flag.Parse()
	err := app.Run()
	assert.Equal(t, err, nil)
	content := readGitignore()
	assert.True(t, strings.Contains(content, "node_modules"))
}

func TestGithubCpp(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	rmGitFiles()
	os.Args = []string{"???", "--name=C++", "--yes", "--verbose"}
	app := NewApp()
	flag.Parse()
	err := app.Run()
	assert.Equal(t, err, nil)
	content := readGitignore()
	assert.True(t, strings.Contains(content, "*.dll"))
}

func TestToptalAngular(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	rmGitFiles()
	os.Args = []string{"???", "--name=Angular", "--yes", "--verbose"}
	app := NewApp()
	flag.Parse()
	err := app.Run()
	assert.Equal(t, err, nil)
	content := readGitignore()
	//fmt.Println(content)
	assert.True(t, strings.Contains(content, "e2e"))
	assert.True(t, strings.Contains(content, "Angular"))
}

func TestFakeFile(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"???", "--name=FakeFile", "--verbose"}
	app := NewApp()
	flag.Parse()
	err := app.Run()
	assert.Equal(t, err, errNoFilesFound)
}
