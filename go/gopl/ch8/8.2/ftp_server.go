package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"
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
			input := strings.TrimSpace(scanner.Text())
			switch {
			case input == "close":
				c.Close()
			case strings.HasPrefix(input, "get"):
				args := strings.Split(input, " ")
				if len(args) < 2 {
					c.Write([]byte("please add a filename"))
				}
				for _, fileName := range args[1:] {
					c.Write([]byte(fmt.Sprintf("download file: %s", fileName)))
				}

			default:
				args := strings.Split(input, " ")
				cmd := exec.Command(args[0], args[1:]...)
				output, err := cmd.CombinedOutput()
				if err != nil {
					c.Write([]byte(err.Error() + "\n"))
				}
				c.Write(output)
			}
		}
	}
}
