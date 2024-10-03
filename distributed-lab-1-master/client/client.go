package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func read(conn *net.Conn) {
	reader := bufio.NewReader(*conn)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from server:", err)
			break
		}
		fmt.Println(":: ", line)
	}
}

func write(conn *net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter text:")
		text, _ := reader.ReadString('\n')
		if text == "/quit\n" {
			fmt.Println("Disconnecting...")
			break
		}
		_, err := fmt.Fprintf(*conn, text)
		if err != nil {
			fmt.Println("Error sending to server:", err)
			break
		}
	}
}

func main() {
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()

	// Try to connect to the server
	conn, err := net.Dial("tcp", *addrPtr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close() // 연결 종료 시 닫기

	// Start asynchronously reading and displaying messages
	go read(&conn)

	// Start getting and sending user messages
	write(&conn)

	fmt.Println("Connection closed.")
}
