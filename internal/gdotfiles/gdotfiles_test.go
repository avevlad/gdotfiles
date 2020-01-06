package gdotfiles

import (
	"flag"
	"github.com/avevlad/gdotfiles/internal/test/assert"
	"os"
	"testing"
)

func TestGithubNode(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"???", "--name=Node", "--yes", "--verbose"}
	app := NewApp()

	flag.Parse()
	err := app.Run()

	assert.Equal(t, err, nil)
}
