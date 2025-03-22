package main

import (
	"fmt"
	"httpfromtcp/internal/request"
	"log"
	"net"
)

const port = ":42069"

// func getLinesChannel(conn io.ReadCloser) <-chan string {
// 	lines := make(chan string)
// 	go func() {
// 		defer close(lines)
// 		defer conn.Close()
// 		currentLineContents := ""
// 		for {
// 			b := make([]byte, 8)
// 			n, err := conn.Read(b)
// 			if err != nil {
// 				if currentLineContents != "" {
// 					lines <- currentLineContents
// 				}
// 				if errors.Is(err, io.EOF) {
// 					break
// 				}
// 				fmt.Printf("error: %s\n", err.Error())
// 				return
// 			}
// 			str := string(b[:n])
// 			parts := strings.Split(str, "\n")
// 			for i := 0; i < len(parts)-1; i++ {
// 				lines <- fmt.Sprintf("%s%s", currentLineContents, parts[i])
// 				currentLineContents = ""
// 			}
// 			currentLineContents += parts[len(parts)-1]
// 		}
// 	}()
// 	return lines
// }

func main() {
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("y u no listen: %s", err)
	}
	defer l.Close()

	fmt.Println("i listen to", port)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("y u no assep: %s", err)
		}

		fmt.Printf("listen assep: %v\n", conn.LocalAddr())
		fmt.Println("=====================================")

		request, err := request.RequestFromReader(conn)
		if err != nil {
			log.Fatalf("Request bad: %s", err)
		}

		fmt.Printf("Request line:\n- Method: %s\n- Target: %s\n- Version: %s\n",
			request.RequestLine.Method, request.RequestLine.RequestTarget, request.RequestLine.HttpVersion)

	}
}
