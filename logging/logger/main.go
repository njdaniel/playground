package main

import (
	"log"
	"os"
	"time"
)

var (
	Error *log.Logger
	Debug *log.Logger
	Info  *log.Logger
)

func init() {
	f := log.LstdFlags | log.Lshortfile
	Error = log.New(os.Stderr, "Error ", f)
	Debug = log.New(os.Stdout, "Debug ", f)
	Info = log.New(os.Stdout, "Info ", f)
}

func main() {
	timeNow := time.Now()
	Info.Println(timeNow)

}
