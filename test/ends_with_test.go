package test

import (
	gotools "github.com/CrestFallenTurtle/Go-tools"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndsWith(t *testing.T) {
	result := gotools.EndsWith("Hello, world!", []string{"!"})

	ok := result["!"]

	assert.Equal(t, true, ok)
}

func TestDoesntEndWith(t *testing.T) {
	result := gotools.EndsWith("Hello, world!", []string{"H"})

	ok := result["H"]

	assert.Equal(t, false, ok)
}
