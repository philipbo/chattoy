package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("---> done")
		done <- struct{}{}
	}()

   go mustCopy(conn, os.Stdin)

	log.Println(1111)
	<- done // wait for background goroutine to finish
	conn.Close()
	log.Println("*****done*****")
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
