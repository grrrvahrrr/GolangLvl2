//Lesson 9 Homework
package main

import (
	"GolangLvl2/lesson2/config"
	"fmt"
	"os"
	"strings"
)

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

func main() {
	var cfg config.Configuration
	var err error

	fmt.Println(`Please input config file name or "d" for default`)
	var filename string
	if _, err = fmt.Scan(&filename); err != nil {
		fmt.Printf("Invalid filename: %v", err)
		os.Exit(1)
	}
	if filename != "d" {
		cfg, err = config.Load(filename)
		if err != nil {
			fmt.Printf("Couldn't load config %s", err)
		}
	} else if filename == "d" {
		cfg, err = config.Load("config_example.env")
		if err != nil {
			fmt.Printf("Couldn't load config %s", err)
		}
	}

	if strings.Contains("hello", cfg.Mode) {
		Greeter(cfg.ParseNames())
	}
	if strings.Contains("gb", cfg.Mode) {
		Goodbyer(cfg.ParseNames())
	}
}
