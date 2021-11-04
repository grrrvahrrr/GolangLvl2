package doc

import (
	"GolangLvl2/lesson2/config"
	"fmt"
	"strings"
)

func ExampleGreeter() {
	var cfg config.Configuration
	var err error

	cfg, err = config.Load("config.json")
	if err != nil {
		fmt.Printf("Couldn't load config %s", err)
	}

	if strings.Contains("hello", cfg.Mode) {
		Greeter(cfg.ParseNames())
	}
	// Output:
	// | Hello | JsonName |
	// | Hello | JsonName2 |
}

func ExampleGoodbyer() {
	var cfg config.Configuration
	var err error

	cfg, err = config.Load("config.yaml")
	if err != nil {
		fmt.Printf("Couldn't load config %s", err)
	}

	if strings.Contains("gb", cfg.Mode) {
		Goodbyer(cfg.ParseNames())
	}
	//Output:
	//| Goodbye | yamlName |
	//| Goodbye | yamlName2 |
}
