// The program is used for greeting or saying goodbye to a user or multiple users in a terminal
//
// Function Greeter accepts a slice of names and prints greetings
// Greeter(names []string)
//
// Function Goodbyer accepts a slice of names and prints goodbye
// Goodbyer(names []string)
//
// Names and Mode ("hello" or "gb") can be provided with flags:
// --name=John,Sam,Niel
// --mode=hello
//
// Also .yaml, .json can be created for setting names and mode. See examples inside the package
// Package config processes the program's configuration through flags and config files
//
// functions Load checks the type of config provided(flag, environment variables, .json, .yaml) and returns type Configuration and an error
// : func Load(file string) (Configuration, error)
//
// Method ParseNames processes names that were input and separates them and returns a slice, that can be passed into Greeter or Goodbyes functions
// : func (c *Configuration) ParseNames() []string
//
//
package doc

import "fmt"

func Greeter(names []string) {
	for i := 0; i < len(names); i++ {
		fmt.Printf("| %s | %s |\n", "Hello", names[i])
	}
	return
}

func Goodbyer(names []string) {
	for i := 0; i < len(names); i++ {
		fmt.Printf("| %s | %s |\n", "Goodbye", names[i])
	}
	return
}
