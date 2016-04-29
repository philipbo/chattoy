package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	port := flag.String("p", ":3000", "chat server port")
	flag.Parse()

	conn, err := net.Dial("tcp", *port)
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		done <- struct{}{}
	}()

	go mustCopy(conn, os.Stdin)

	<-done // wait for background goroutine to finish
	conn.Close()
	log.Println("*****Quit*****")
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
