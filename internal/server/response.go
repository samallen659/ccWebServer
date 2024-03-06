package server

import (
	"bytes"
	"errors"
	"fmt"

	"golang.org/x/exp/slices"
)

type HTTPStatus int

const (
	HTTP_OK                   HTTPStatus = 200
	HTTP_BAD_REQUEST          HTTPStatus = 400
	HTTP_NOT_FOUND            HTTPStatus = 404
	HTTP_INTERNAL_SEVER_ERROR HTTPStatus = 500
)

var HTTPStatuses = []HTTPStatus{HTTP_OK, HTTP_NOT_FOUND, HTTP_INTERNAL_SEVER_ERROR}

type ResponseHeader struct {
	Status     HTTPStatus
	Additional map[string]string
}

type Response struct {
	Header *ResponseHeader
	Body   string
}

func NewResponse() *Response {
	return &Response{
		&ResponseHeader{},
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
