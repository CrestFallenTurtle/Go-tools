package gotools

import "path/filepath"

// Returns the executable name
func GrabExecutableName() string {
	path := GrabExecutablePath()
	exe_name := filepath.Base(path)

	return exe_name
}
