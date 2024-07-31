package gotools

import "os"

// Returns the path to the executable file
// including the executable file name
func GrabExecutablePath() string {
	path := os.Args[0]
	return path
}
