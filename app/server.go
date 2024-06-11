package main

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/pkg/handle"
	"net"
	"os"
	"path"
	"strconv"
	"strings"
)

func handler(conn net.Conn) error {
	request, err := handle.HTTPRequest(conn)
	if err != nil {
		return err
	}

	switch {
	case request.URL == "/":
		err = handle.RespondWithCode(conn, 200)
	case strings.HasPrefix(request.URL, "/echo/"):
		bodyString := request.URL[len("/echo/"):]
		headers := map[string]string{
			"Content-Type":   "text/plain",
			"Content-Length": strconv.Itoa(len(bodyString)),
		}

		if encoding, ok := request.Headers["Accept-Encoding"]; ok && encoding == "gzip" {
			headers["Accept-Encoding"] = encoding
		}

		err = handle.Respond(
			conn,
			200,
			headers,
			[]byte(bodyString),
		)
	case strings.HasPrefix(request.URL, "/files/"):
		switch request.Method {
		case "GET":
			dirPath := os.Args[2]
			var data []byte
			data, err = os.ReadFile(path.Join(dirPath, request.URL[len("/files/"):]))
			if err != nil {
				err = handle.RespondWithCode(conn, 404)
				break
			}

			err = handle.Respond(
				conn,
				200,
				map[string]string{
					"Content-Type":   "application/octet-stream",
					"Content-Length": strconv.Itoa(len(data)),
				},
				data,
			)
		case "POST":
			dirPath := os.Args[2]
			err = os.WriteFile(
				path.Join(dirPath, request.URL[len("/files/"):]),
				request.Body,
				0644,
			)
			if err != nil {
				break
			}

			err = handle.RespondWithCode(conn, 201)

		default:
			err = handle.RespondWithCode(conn, 405)
		}

	case request.URL == "/user-agent":
		bodyString := request.Headers["User-Agent"]
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
	return err
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go func() {
			err := handler(conn)
			if err != nil {
				fmt.Println("Error handling HTTP request: ", err.Error())
				os.Exit(3)
			}
		}()
	}

}
