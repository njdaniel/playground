package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println(time.Now().Format("2006-01-02T15:04:05"))
	files := []string{"file1.log"}
	AggregateLogs(files)
}

// print out multiple log records from multiple log files in order

//example
//timestamp msg
//EDIT: log format to 2006-01-02T15:04:05
//file1.log
//2020-03-25T15:18:26 msg 1
//2020-03-25T15:18:26 msg 2
//2020-03-25T15:18:26 msg 3
//2020-03-25T15:18:27 msg 4

//file2.log
//2020-03-25T15:18:26 msg 5
//2020-03-25T15:18:26 msg 6
//2020-03-26T15:18:21 msg 7
//2020-03-26T15:18:26 msg 8

//file3.log
//2020-03-25T14:18:26 msg 9
//2020-03-25T16:18:26 msg 10
//2020-03-25T17:18:26 msg 11
//2020-03-25T17:19:26 msg 12

//Result prints
// TODO: check if results are correct
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
	Msg string
}

// AggregateLogs prints out the log records in order
func AggregateLogs(files []string) {
	ch := make(chan time.Time)
	//done := make(chan bool)
	//logTimes := make([]chan time.Time, len(files))
	for _, file := range files {
		//setup goroutine
		go func() {
			readFile, err := os.Open(file)
			if err != nil {
				log.Fatalf("error failed to open file: %v", err)
			}
			fileScanner := bufio.NewScanner(readFile)
			fileScanner.Split(bufio.ScanLines)
			for fileScanner.Scan() {
				tss := strings.Split(fileScanner.Text(), " ")[0]
				//fmt.Println(ts)
				layout := "2006-01-02T15:04:05"
				ts, err := time.Parse(layout, tss)
				if err != err {
					fmt.Println(err)
				}
				ch <- ts
			}
			close(ch)
			//<-done
		}()
		for {
			ts, ok := <- ch
			if !ok {
				break
			}
			fmt.Println(ts.String())
		}
		//<-done
		//get timestamp
		//send back
	}
	//compare all the returned timestamps
	//for ;len(logTimes) == 0; {
	//	for k, v := range logTimes {
	//		var oldest time.Time
	//
	//	}
	//}
}

