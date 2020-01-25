package main

import (
	"flag"
	"fmt"
)

func main() {
	argConfig := flag.String("config", "./boomi_automate.yml", "Path to config file")
	argSecret := flag.String("secret", "", "Path to config file")
	flag.Parse()
	fmt.Println("Config:", *argConfig)
	fmt.Println("Secret:", *argSecret)
}
