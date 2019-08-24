package main

import (
	"io"
	"net"
	"log"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
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
	conn.Close()
	<-done // wait for background goroutine to finish
	log.Println("done")
}

func mustCopy(dst io.Writer, src io.Reader)  {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
