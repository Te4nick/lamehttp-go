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
