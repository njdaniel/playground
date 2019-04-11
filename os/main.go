package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("bash", "-c", "git clone https://github.com/njdaniel/test.git ~/tmp/test")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	cmd.Wait()
	// if err != nil {
	// 	fmt.Printf("could not execute: %v", err)
	// }
	// fmt.Println(string(out[:]))
}
