package handle

import (
	"bytes"
	"compress/gzip"
	lameHTTP "github.com/codecrafters-io/http-server-starter-go/app/pkg/lamehttp"
	"net"
	"strconv"
	"strings"
)

type HandlerFunc func(request *lameHTTP.Request) (*lameHTTP.Response, error)

type Endpoint struct {
	pathTrie *URITrie
}

func NewEndpoint() *Endpoint {
	return &Endpoint{
		pathTrie: NewURITrie(),
	}
}

func (e *Endpoint) AddPath(path string, method []string, handler HandlerFunc) {
	if path[0] == '/' {
		path = path[1:]
	}

	methods := make(map[string]HandlerFunc)
	for _, m := range method {
		methods[strings.ToLower(m)] = handler
	}

	before, _, _ := strings.Cut(path, "?")
	e.pathTrie.Put(before, methods)
}

func (e *Endpoint) HandleConnection(conn net.Conn) error {
	request, err := HTTPRequest(conn)
	if err != nil {
		return RespondWithCode(conn, 400)
	}

	path := request.URL
	if queryIdx := strings.IndexRune(request.URL, '?'); queryIdx != -1 {
		path = path[:queryIdx]
	}

	uri := e.pathTrie.Get(path)
	if uri == nil {
		return RespondWithCode(conn, 404)
	}

	method, ok := uri.Methods[strings.ToLower(request.Method)]
	if !ok {
		return RespondWithCode(conn, 405)
	}

	resp, err := method(request)
	if err != nil {
		connErr := RespondWithCode(conn, 500)
		if connErr != nil {
			return connErr
		}
	}

	if resp == nil {
		return RespondWithCode(conn, 404) // TODO: what to return if no response?
	}

	if resp.Body == nil {
		return Respond(conn, resp.Status, resp.Headers, resp.Body)
	}

	if encoding, ok := request.Headers["Accept-Encoding"]; ok && strings.Contains(encoding, "gzip") {
		var buf bytes.Buffer
		zw := gzip.NewWriter(&buf)
		_, gzipErr := zw.Write(resp.Body)
		if gzipErr != nil {
			err = gzipErr
		}
		gzipErr = zw.Close()
		if gzipErr != nil {
			err = gzipErr
		}
		resp.Body = buf.Bytes()

		resp.Headers["Content-Length"] = strconv.Itoa(len(buf.String()))
		resp.Headers["Content-Encoding"] = "gzip"

		return Respond(conn, resp.Status, resp.Headers, resp.Body)
	}
	return err
}
