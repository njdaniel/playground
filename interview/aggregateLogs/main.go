package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	fmt.Println(time.Now().Format("2006-01-02T15:04:05"))
	files := []string{"file1.log", "file2.log"}
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
var wg sync.WaitGroup
// AggregateLogs prints out the log records in order
func AggregateLogs(files []string) {
	wg.Add(2)
	ch1 := make(chan time.Time)
	ch2 := make(chan time.Time)
	//done := make(chan bool)
	//logTimes := make([]chan time.Time, len(files))
	go sendTimeStamps(files[0], ch1)
	go sendTimeStamps(files[1], ch2)

	for {
		// TODO: can only compare the channel if it == to the same ref data structure or nil
		if ch1 < ch2 {
			ts, ok := <- ch1
			if !ok {
				break
			}
			fmt.Println(ts.String())
		} else if ch2 < ch1 {
			ts, ok := <- ch2
			if !ok {
				break
			}
			fmt.Println(ts.String())
		} else if ch2 == ch1 {
			ts1, ok := <- ch1
			if !ok {
				break
			}
			ts2, ok := <- ch2
			if !ok {
				break
			}
			fmt.Println(ts1.String())
			fmt.Println(ts2.String())
		}
	}
	wg.Wait()
	// TODO: compare all the returned timestamps
	//for ;len(logTimes) == 0; {
	//	for k, v := range logTimes {
	//		var oldest time.Time
	//
	//	}
	//}
}

func sendTimeStamps(file string, ch chan time.Time)  {
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
	wg.Done()
	//<-done
}

