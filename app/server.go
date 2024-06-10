package main

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/pkg/handle"
	lamehttp "github.com/codecrafters-io/http-server-starter-go/app/pkg/lamehttp"
	"net"
	"os"
	"strconv"
	"strings"
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
	//_, err = conn.Write((&lame_http.Response{
	//	Status: 200,
	//}).Byte())
	//if err != nil {
	//	fmt.Println("Error writing to connection: ", err.Error())
	//	os.Exit(2)
	//}
	bytes := make([]byte, 1024)
	_, err = conn.Read(bytes)
	if err != nil {
		fmt.Println("Error reading from connection: ", err.Error())
		os.Exit(2)
	}
	request, err := lamehttp.ParseHTTPRequest(bytes)
	if err != nil {
		fmt.Println("Error parsing request: ", err.Error())
		os.Exit(4)
	}
	switch {
	case request.URL == "/":
		err = handle.RespondWithCode(conn, 200)
	case strings.HasPrefix(request.URL, "/echo/"):
		bodyString := request.URL[len("/echo/"):]
		err = handle.Respond(
			conn,
			200,
			map[string]string{
				"Content-Type":   "text/plain",
				"Content-Length": strconv.Itoa(len(bodyString)),
			},
			[]byte(bodyString),
		)
	default:
		err = handle.RespondWithCode(conn, 404)
	}

	if err != nil {
		fmt.Println("Error during responding: ", err.Error())
		os.Exit(3)
	}
}
