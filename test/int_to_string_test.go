package test

import (
	gotools "github.com/CrestFallenTurtle/Go-tools"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntToString(t *testing.T) {
	result := gotools.IntToString(666)

	assert.Equal(t, "666", result)
}
