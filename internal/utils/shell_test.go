package utils

import (
	"testing"

	"github.com/avevlad/gdotfiles/internal/test/assert"
)

func TestRealGit(t *testing.T) {
	assert.True(t, CheckBinExist("git", "--version"))
}
func TestFzfExist(t *testing.T) {
	assert.True(t, CheckFzfExist())
}
