package logging

import (
	"errors"
	"testing"
)

func TestError(t *testing.T) {
	Error(errors.New("Warning!"), "Error at level %s and %s", "three", "four")
}
