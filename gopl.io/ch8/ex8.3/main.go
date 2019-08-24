package main

import (
	"io"
	"log"
	"net"
	"os"
)

// modify netcat3 to just close the tcp write connection and not the read connection
func main() {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring error
		log.Println("done")
		done <- struct{}{} // signal to main goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.CloseWrite()
	<-done // wait for background goroutine to finish
}
func mustCopy(dst io.Writer, src io.Reader)  {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

