package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

const port = ":42069"

func getLinesChannel(conn io.ReadCloser) <-chan string {
	lines := make(chan string)
	go func() {
		defer close(lines)
		defer conn.Close()
		currentLineContents := ""
		for {
			b := make([]byte, 8)
			n, err := conn.Read(b)
			if err != nil {
				if currentLineContents != "" {
					lines <- currentLineContents
				}
				if errors.Is(err, io.EOF) {
					break
				}
				fmt.Printf("error: %s\n", err.Error())
				return
			}
			str := string(b[:n])
			parts := strings.Split(str, "\n")
			for i := 0; i < len(parts)-1; i++ {
				lines <- fmt.Sprintf("%s%s", currentLineContents, parts[i])
				currentLineContents = ""
			}
			currentLineContents += parts[len(parts)-1]
		}
	}()
	return lines
}

func main() {
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("y u no listen: %s", err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("y u no assep: %s", err)
		}

		fmt.Printf("listen assep: %v\n", conn.LocalAddr())
		fmt.Println("=====================================")

		outChan := getLinesChannel(conn)

		for line := range outChan {
			fmt.Printf("%s\n", line)
		}
		fmt.Println("connection to ", conn.RemoteAddr(), "yeeted")
	}

}
