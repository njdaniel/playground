package main

import (
	"net"
	"log"
	"bufio"
	"time"
	"fmt"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:4000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn)  {
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}

func echo(c net.Conn, shout string, delay time.Duration)  {
	fmt.Println(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}