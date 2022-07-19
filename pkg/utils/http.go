package utils

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

func GetClient() *http.Client {
	return &http.Client{
		Timeout: 2 * time.Second,
	}
}

func ResponseHttpBody(bodyIO io.ReadCloser) (body []byte) {
	buffer := bytes.NewBuffer(make([]byte, 0, 65536))
	io.Copy(buffer, bodyIO)
	temp := buffer.Bytes()
	length := len(temp)
	if cap(temp) > (length + length/10) {
		body = make([]byte, length)
		copy(body, temp)
	} else {
		body = temp
	}
	buffer.Reset()
	return body
}
