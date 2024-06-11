package lameHTTP

import "strconv"

const HTTPVersion string = "1.1"

var HTTPStatusString = map[int]string{
	200: "OK",
	201: "Created",
	404: "Not Found",
	405: "Method Not Allowed",
}

type Response struct {
	Status  int
	Headers map[string]string
	Body    []byte
}

func (r *Response) Byte() []byte {
	versionLine := "HTTP/" + HTTPVersion
	var statusLine string
	if statusString, ok := HTTPStatusString[r.Status]; ok {
		statusLine = versionLine + " " + strconv.Itoa(r.Status) + " " + statusString
	} else {
		statusLine = versionLine + " " + strconv.Itoa(r.Status)
	}

	headersLine := ""
	for key, value := range r.Headers {
		headersLine += key + ": " + value + "\r\n"
	}
	return append([]byte(statusLine+"\r\n"+headersLine+"\r\n"), r.Body...)
}
