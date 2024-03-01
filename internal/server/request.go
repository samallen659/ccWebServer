package server

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/exp/slices"
)

var PATH_REGEX = `^\/(?!.*\/\/)([a-zA-Z-\/]+)$`

type HTTPMethod string

const (
	HTTP_GET    HTTPMethod = "GET"
	HTTP_POST   HTTPMethod = "POST"
	HTTP_HEAD   HTTPMethod = "HEAD"
	HTTP_DELETE HTTPMethod = "DELETE"
	HTTP_PATCH  HTTPMethod = "PATCH"
)

var HTTP_METHODS = []HTTPMethod{HTTP_GET, HTTP_POST, HTTP_HEAD, HTTP_DELETE, HTTP_PATCH}

type Header struct {
	Method     HTTPMethod
	Path       string
	Additional map[string]string
}

type Request struct {
	Header *Header
	Body   string
}

func NewHeader(headStr string) (*Header, error) {
	lines := strings.Split(headStr, "\n")
	fLine := strings.Split(lines[0], " ")

	method := HTTPMethod(fLine[0])
	if !slices.Contains(HTTP_METHODS, method) {
		return nil, errors.New("Invalid HTTP Method")
	}

	validPath := regexp.MustCompile(PATH_REGEX)
	path := fLine[1]
	if !validPath.Match([]byte(path)) {
		return nil, errors.New("Invalid URL Path")
	}

	fields := make(map[string]string)
	for i := 1; i < len(lines); i++ {
		field := strings.Split(lines[i], ":")
		if len(field) != 2 {
			return nil, errors.New("Invalid Field")
		}
		fields[field[0]] = field[1]
	}

	return &Header{
		Method:     method,
		Path:       path,
		Additional: fields,
	}, nil
}

func NewRequest(reqStr string) (*Request, error) {
	components := strings.Split(reqStr, "\r\n\r\n")
	if len(components) < 2 {
		return nil, errors.New("Invalid HTTP Request")
	}

	header, err := NewHeader(components[0])
	if err != nil {
		return nil, fmt.Errorf("Error creating header: %s", err.Error())
	}

	return &Request{
		Header: header,
		Body:   components[1],
	}, nil
}
