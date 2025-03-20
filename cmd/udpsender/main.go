package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const prompt = ">"

func main() {
	adr, err := net.ResolveUDPAddr("udp", "127.0.0.1:42069")
	if err != nil {
		log.Fatalf("UDP Ungood: %s", err)
	}

	conn, err := net.DialUDP("udp", nil, adr)
	if err != nil {
		log.Fatalf("conn ungood: %s", err)
	}

	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s", prompt)
		line, err := r.ReadString('\n')
		if err != nil {
			log.Fatalf("No read: %s", err)
		}

		_, err = conn.Write([]byte(line))
		if err != nil {
			log.Printf("No write: %s", err)
		}
	}
}
