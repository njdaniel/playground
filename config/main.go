package main

import (
	"github.com/kylelemons/go-gypsy/yaml"
	"fmt"
)

func main() {
	config, err := yaml.ReadFile("config.yaml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T \n", config)
	fmt.Printf("Path is type: %T", (config.Get("path")))
	fmt.Println(config.GetBool("enabled"))
	fmt.Println(config.Get("path"))
	fmt.Println(config.GetInt("count"))
}
