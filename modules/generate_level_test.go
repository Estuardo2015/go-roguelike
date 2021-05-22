package modules

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateLevelFromTxt(t *testing.T) {
	path := "./assets/levels/unnamed.txt"
	lvl, err := CreateLevelFromTxtFile(path)
	assert.Nil(t, err)
	assert.NotNil(t, lvl)
}
