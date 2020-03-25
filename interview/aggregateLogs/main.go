package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println(time.Now())
}

// print out multiple log records from multiple log files in order

//example
//timestamp msg
//file1.log
//2020-03-25 15:18:26.04957 msg 1
//2020-03-25 15:18:26.04957 msg 2
//2020-03-25 15:18:26.04957 msg 3
//2020-03-25 15:18:27.04957 msg 4

//file2.log
//2020-03-25 15:18:26.03957 msg 5
//2020-03-25 15:18:26.14957 msg 6
//2020-03-26 15:18:21.04957 msg 7
//2020-03-26 15:18:26.04958 msg 8

//file3.log
//2020-03-25 14:18:26.04957 msg 9
//2020-03-25 16:18:26.04957 msg 10
//2020-03-25 17:18:26.04957 msg 11
//2020-03-25 17:19:26.04957 msg 12

//Result prints
// 9
// 5
// 1
// 2
// 3
// 6
// 7
// 8
// 4
// 10
// 11
// 12

// log files can be very large
// assume each files have the logs in order

//open each file in goroutines
// pass the value back to the main function and print the oldest value
// get new value

type Log struct {
	TimeStamp time.Time
	msg string
}

func AggregateLogs(files []string) {
	logTimes := make([]chan time.Time, len(files))
	for _, file := range files {
		readFile, err := os.Open(file)
		if err != nil {
			log.Fatalf("error failed to open file: %v", err)
		}
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
	}
}
