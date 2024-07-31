package gotools

import (
	"fmt"

	notify "github.com/CrestFallenTurtle/notify_handler"
)

// Converts the provided string to into a boolean
func StringToBoolean(value string) bool {
	to_return := false

	if value != "true" && value != "false" {
		notify.Error(fmt.Sprintf("Needed true/false, recieved %s", value), "gotools.StringToBoolean()", 1)
	}

	if value == "true" {
		to_return = true
	}

	return to_return
}
