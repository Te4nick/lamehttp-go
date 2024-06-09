package main

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/pkg/response"
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	_, err = conn.Write((&response.HTTP{
		Status: 200,
	}).Byte())
	if err != nil {
		fmt.Println("Error writing to connection: ", err.Error())
		os.Exit(2)
	}

	err = conn.Close()
	if err != nil {
		fmt.Println("Error closing the connection: ", err.Error())
		os.Exit(3)
	}
}
