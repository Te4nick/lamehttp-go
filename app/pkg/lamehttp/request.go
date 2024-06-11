package lameHTTP

import (
	"bufio"
	"bytes"
	"errors"
	"strconv"
	"strings"
)

type Request struct {
	Method  string
	URL     string
	Version string
	Headers map[string]string
	Body    []byte
}

func ParseHTTPRequest(data []byte) (*Request, error) {
	req := &Request{}
	br := bytes.NewReader(data)
	buf := bufio.NewReader(br)
	line, err := buf.ReadString('\n')
	if err != nil {
		return nil, err
	}
	parts := strings.SplitN(line, " ", 3)
	if len(parts) != 3 {
		return nil, errors.New("invalid HTTP request")
	}
	req.Method = parts[0]
	req.URL = parts[1]
	req.Version = parts[2]
	headers := make(map[string]string)
	for {
		line, err = buf.ReadString('\n')
		if err != nil {
			break
		}
		if line == "\r\n" {
			break
		}
		parts = strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			return nil, errors.New("invalid HTTP header")
		}
		headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
	}
	req.Headers = headers

	body := make([]byte, 1024)
	_, err = buf.Read(body)
	if err != nil {
		return nil, err
	}

	if body[0] != '\r' && body[1] != '\n' {
		contentLength, err := strconv.Atoi(headers["Content-Length"])
		if err != nil {
			return nil, err
		}

		req.Body = body[:contentLength]
	}

	return req, nil
}
