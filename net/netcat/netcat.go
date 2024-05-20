package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var (
	port = flag.Int("p", 3090, "port to use")
	host = flag.String("h", "localhost", "host to use")
)

func main() {
	flag.Parse()
	// connect to the server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatalf("Connection error:%v\n", err)
	}
	// create channel to handle communication connection
	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, conn)
		done <- struct{}{}
	}()
	CopyContent(conn, os.Stdin)
	conn.Close()
	<-done
}

func CopyContent(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Fatalf("Copying error: %v\n", err)
	}
}
