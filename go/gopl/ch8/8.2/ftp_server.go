package main

import (
	"bufio"
	"flag"
	"log"
	"net"
)

func main() {
	port := flag.String("port", "8000", "listen port")
	flag.Parse()
	listener, err := net.Listen("tcp", "localhost:"+*port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		scanner := bufio.NewScanner(c)
		for scanner.Scan() {
			c.Write([]byte(scanner.Text()))
		}
	}
}
