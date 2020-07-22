package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func handleConn(c net.Conn, location *time.Location) {
	defer c.Close()
	for {
		now := time.Now()
		_, err := io.WriteString(c, now.In(location).Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	port := flag.String("port", "8000", "listen port")
	zone := flag.String("zone", "Asia/Shanghai", "timezone")
	flag.Parse()
	location, err := time.LoadLocation(*zone)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening to", *port)
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
		go handleConn(conn, location) // handle connections concurrently
	}
}
