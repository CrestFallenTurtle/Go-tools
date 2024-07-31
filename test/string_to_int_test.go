package test

import (
	gotools "github.com/CrestFallenTurtle/Go-tools"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringToInt(t *testing.T) {
	result := gotools.StringToInt("666")

	assert.Equal(t, 666, result)
}

func TestStringToIntInvalid(t *testing.T) {
	result := gotools.StringToInt("true")

	assert.Equal(t, -1, result)
}
