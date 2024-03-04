package server

import (
	"bytes"
	"errors"
	"fmt"

	"golang.org/x/exp/slices"
)

type HTTPStatus int

const (
	HTTP_OK HTTPStatus = 200
)

var HTTPStatuses = []HTTPStatus{HTTP_OK}

type ResponseHeader struct {
	Status     HTTPStatus
	Additional map[string]string
}

type Response struct {
	Header ResponseHeader
	Body   string
}

func NewResponse() *Response {
	return &Response{
		ResponseHeader{},
		"",
	}
}

func (r *Response) SetStatus(status HTTPStatus) error {
	if !slices.Contains(HTTPStatuses, status) {
		return errors.New("Invalid HTTP Status")
	}

	r.Header.Status = status
	return nil
}

func (r *Response) SetHeader(key string, value string) {
	r.Header.Additional[key] = value
}

func (r *Response) SetBody(body string) {
	r.Body = body
}

func (r *Response) Marshall() []byte {
	var b bytes.Buffer
	b.Write([]byte(fmt.Sprintf("HTTP/1.1 %d OK\r\n\r\n", r.Header.Status)))
	b.Write([]byte(r.Body))

	return b.Bytes()
}
