package response

import "strconv"

const HTTPVersion string = "1.1"

var HTTPStatusString = map[int]string{
	200: "OK",
}

type HTTP struct {
	Status  int
	Headers map[string]string
	Body    []byte
}

func (r *HTTP) Byte() []byte {
	versionLine := "HTTP/" + HTTPVersion
	var statusLine string
	if statusString, ok := HTTPStatusString[r.Status]; ok {
		statusLine = versionLine + " " + strconv.Itoa(r.Status) + " " + statusString
	} else {
		statusLine = versionLine + " " + strconv.Itoa(r.Status)
	}

	headersLine := ""
	bodyLine := ""
	return []byte(statusLine + "\r\n" + headersLine + "\r\n" + bodyLine)
}
