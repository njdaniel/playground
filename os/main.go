package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("bash", "-c", "ls -lah").CombinedOutput()
	if err != nil {
		fmt.Printf("could not execute: %v", err)
	}
	fmt.Println(string(out[:]))
}
