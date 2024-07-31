package gotools

import (
	"math/rand"
	"time"
)

// Generates a random number between min and max
func GenerateRandomIntBetween(min int, max int) int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(max-min) + min
}

// Generates a random number between 1 and 128
func GenerateRandomInt() int {
	return GenerateRandomIntBetween(1, 128)
}
