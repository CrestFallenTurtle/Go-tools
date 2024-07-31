package gotools

// Generates a random bool
func GenerateRandomBool() bool {
	value := GenerateRandomIntBetween(0, 2)
	return value == 1
}
