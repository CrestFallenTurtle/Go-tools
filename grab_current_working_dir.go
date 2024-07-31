package gotools

import (
	"os"

	notify "github.com/CrestFallenTurtle/notify_handler"
)

// Grabs the current working path
func GrabCWD() string {
	path, err := os.Getwd()
	if err != nil {
		notify.Error(err.Error(), "gotools.GrabCWD()", 1)
	}
	return path
}
