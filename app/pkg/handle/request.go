package handle

import (
	lameHTTP "github.com/codecrafters-io/http-server-starter-go/app/pkg/lamehttp"
	"net"
)

func HTTPRequest(conn net.Conn) (*lameHTTP.Request, error) {
	bytes := make([]byte, 1024)
	_, err := conn.Read(bytes)
	if err != nil {
		return nil, err
	}

	request, err := lameHTTP.ParseHTTPRequest(bytes)
	if err != nil {
		return nil, err
	}

	return request, nil
}
