package gotools

import (
	"os/user"

	notify "github.com/CrestFallenTurtle/notify_handler"
)

// Grabs the current users home directory
func GrabHomeDir() string {
	path, err := user.Current()

	if err != nil {
		notify.Error(err.Error(), "gotools.GrabHomeDir()", 1)
	}

	return path.HomeDir
}
