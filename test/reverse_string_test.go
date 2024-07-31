package test

import (
	gotools "github.com/CrestFallenTurtle/Go-tools"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseString(t *testing.T) {
	value := "mephisto"
	gotools.ReverseString(&value)

	assert.Equal(t, "otsihpem", value)
}
