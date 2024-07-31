# Go-tools
The functions found here are tools that I find useful for everyday development. 

## Installation
Simply run `go get github.com/CrestFallenTurtle/Go-tools` to add it to your project

## Contents
The following table contains all tools available in this project.

| Function         | Description     | Note |
|--------------|-----------|-----------|
| `Contains` | Checks if the the supplied string contains a substring | Returns a map[string]bool |
| `EraseDelimiter`| Erases any delimiter from your string |  |
| `GenerateRandomBool`| Generates and returns true or false |  |
| `GenerateRandomIntBetween`| Generates a random int in an interval |  |
| `GenerateRandomInt`| Generates a random int between 1 and 128 |  |
| `GrabCWD`| Returns the current working directory | Exits on error |
| `GrabExecutableName`| Returns the executables name |  |
| `GrabExecutablePath`| Returns the path where the executable is housed |  |
| `GrabHomeDir`| Returns the current users home directory | Exits on error |
| `GrabUsername`| Returns the current users username | Exits on error |
| `ReverseString`| Reverses a string |  |
| `SplitString`| Splits a string and returns a rune array |  |
| `StartsWith`| Checks if a string starts with a substring |  |
| `EndsWith` | Checks if a string ends with a substring  | |
| `StringToBoolean`| Converts a string to a boolean | Requires that the string is "true" or "false" |
| `StringToInt`| Converts a string to a int | Returns `-1` when it fails to convert |
| `IntToString`| Converts a int to a string |  |

## Examples
Check out the files under `test/` to see how each function can be utilized