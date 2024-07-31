package test

import (
	gotools "github.com/CrestFallenTurtle/Go-tools"
	"github.com/stretchr/testify/assert"
	"os/user"
	"testing"
)

func TestGrabUsername(t *testing.T) {
	result := gotools.GrabUsername()
	user, _ := user.Current()

	assert.Equal(t, user.Username, result)
}
