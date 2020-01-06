package gdotfiles

import (
	"flag"
	"fmt"
	"github.com/avevlad/gdotfiles/internal/test/assert"
	"github.com/avevlad/gdotfiles/internal/utils"
	"os"
	"strings"
	"testing"
)

func runApp() error {
	app := NewApp()
	flag.Parse()

	return app.Run()
}

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
	err := runApp()
	assert.Equal(t, err, nil)
	content := readGitignore()
	assert.True(t, strings.Contains(content, "node_modules"))
}

func TestGithubCpp(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	rmGitFiles()
	os.Args = []string{"???", "--name=C++", "--yes", "--verbose"}
	err := runApp()
	assert.Equal(t, err, nil)
	content := readGitignore()
	assert.True(t, strings.Contains(content, "*.dll"))
}

func TestToptalAngular(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	rmGitFiles()
	os.Args = []string{"???", "--name=Angular", "--yes", "--verbose"}
	err := runApp()
	assert.Equal(t, err, nil)
	content := readGitignore()
	//fmt.Println(content)
	assert.True(t, strings.Contains(content, "e2e"))
	assert.True(t, strings.Contains(content, "Angular"))
}

func TestGitAttributes(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	rmGitFiles()
	os.Args = []string{"???", "--name=C++", "--type=attributes", "--yes", "--verbose"}
	err := runApp()

	err2 := runApp()
	fmt.Println(err2)
	assert.Equal(t, err, nil)
	content := readGitAttributes()
	fmt.Println(content)
	assert.True(t, strings.Contains(content, "diff"))
	assert.True(t, strings.Contains(content, "cpp"))
}

func TestFakeFile(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"???", "--name=FakeFile", "--verbose"}
	err := runApp()
	assert.Equal(t, err, errNoFilesFound)
}

func TestBuildAppFlagsFromFzfResponse(t *testing.T) {
	resp := "Go.gitignore                           [github_gitignore]"
	af := buildAppFlagsFromFzfResponse(resp)
	assert.Equal(t, af.Name, "Go")
	assert.Equal(t, af.Type, "gitignore")
	assert.Equal(t, af.From, "github_gitignore")
}
