package main

import (
	"log"
	"os"
)

func main() {
	f := log.LstdFlags | log.Lshortfile
	logger := log.New(os.Stdout, "TestLog ", f)
	logger.Println("Program start")
	logger.Panicln("Panic?")
}
