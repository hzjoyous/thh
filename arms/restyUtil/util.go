package restyUtil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

func GetCurlByR(r resty.Response) bytes.Buffer {
	b2 := bytes.Buffer{}
	b2.WriteString(fmt.Sprintf("curl '%v' -X '%v'", r.Request.URL, r.Request.Method))
	for header, headerValue := range r.Request.Header {
		b2.WriteString(fmt.Sprintf(" -H '%v:%v'", header, headerValue[len(headerValue)-1]))
	}
	if r.Request.Body != nil {
		body, _ := json.Marshal(r.Request.Body)
		b2.WriteString(fmt.Sprintf(" --data-raw '%v' ", string(body)))
	}
	if r.Request.FormData != nil {
		for key, value := range r.Request.FormData {
			b2.WriteString(fmt.Sprintf(` --form '%v="%v" `, key, value))
		}
	}
	b2.WriteString(" --compressed --insecure ")
	return b2
}
