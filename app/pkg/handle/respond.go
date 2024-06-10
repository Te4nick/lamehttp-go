package handle

import (
	lamehttp "github.com/codecrafters-io/http-server-starter-go/app/pkg/lamehttp"
	"net"
)

func RespondWithCode(conn net.Conn, code int) error {
	_, err := conn.Write((&lamehttp.Response{
		Status: code,
	}).Byte())
	if err != nil {
		return err
	}
	return conn.Close()
}

func Respond(conn net.Conn, code int, headers map[string]string, body []byte) error {
	_, err := conn.Write((&lamehttp.Response{
		Status:  code,
		Headers: headers,
		Body:    body,
	}).Byte())
	if err != nil {
		return err
	}
	return conn.Close()
}
