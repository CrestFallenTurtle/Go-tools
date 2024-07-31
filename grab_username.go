package gotools

import (
	"os/user"
	"strings"

	notify "github.com/CrestFallenTurtle/notify_handler"
)

// Returns the username of the current user
func GrabUsername() string {
	user, err := user.Current()

	if err != nil {
		notify.Error(err.Error(), "gotools.GrabUsername()", 1)
	}

	to_return := user.Username

	if strings.Contains(to_return, "\\") { // Only occurs on windows so far
		split := strings.Split(to_return, "\\")
		to_return = split[1]
	}

	return to_return
}
