package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("bash", "-c", "docker run ubuntu bash -c 'while sleep 1; do echo 1; done'")
	cmd.Stdout = os.Stdout
	cmd.Start()
	cmd.Wait()
	// if err != nil {
	// 	fmt.Printf("could not execute: %v", err)
	// }
	// fmt.Println(string(out[:]))
}
