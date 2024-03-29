package server

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"

	"golang.org/x/exp/slices"
)

var fileRegex = regexp.MustCompile(`&*([a-zA-Z]*).html`)

type HTTPMethod string

const (
	HTTP_GET    HTTPMethod = "GET"
	HTTP_POST   HTTPMethod = "POST"
	HTTP_HEAD   HTTPMethod = "HEAD"
	HTTP_DELETE HTTPMethod = "DELETE"
	HTTP_PATCH  HTTPMethod = "PATCH"
)

var HTTP_METHODS = []HTTPMethod{HTTP_GET, HTTP_POST, HTTP_HEAD, HTTP_DELETE, HTTP_PATCH}

type RequestHeader struct {
	Method     HTTPMethod
	Path       string
	File       string
	Additional map[string]string
}

type Request struct {
	Header *RequestHeader
	Body   string
}

func NewRequestHeader(headStr string) (*RequestHeader, error) {
	lines := strings.Split(headStr, "\n")
	fLine := strings.Split(lines[0], " ")

	method := HTTPMethod(fLine[0])
	if !slices.Contains(HTTP_METHODS, method) {
		return nil, errors.New("Invalid HTTP Method")
	}

	path := fLine[1]
	file := ""
	if fileRegex.Match([]byte(path)) {
		log.Println("Requested path contains file")
		path, file = parsePathAndFile(path)
		log.Printf("Parsed path to %s and file %s", path, file)
	}

	path = strings.TrimSuffix(path, "/")

	fields := make(map[string]string)
	for i := 1; i < len(lines); i++ {
		field := strings.Split(lines[i], ": ")
		if len(field) != 2 {
			return nil, errors.New("Invalid Field")
		}
		fields[field[0]] = field[1]
	}

	return &RequestHeader{
		Method:     method,
		Path:       path,
		File:       file,
		Additional: fields,
	}, nil
}

func NewRequest(reqStr string) (*Request, error) {
	components := strings.Split(reqStr, "\r\n\r\n")
	if len(components) < 2 {
		return nil, errors.New("Invalid HTTP Request")
	}

	header, err := NewRequestHeader(components[0])
	if err != nil {
		return nil, fmt.Errorf("Error creating header: %s", err.Error())
	}

	return &Request{
		Header: header,
		Body:   components[1],
	}, nil
}

func parsePathAndFile(p string) (string, string) {
	i := fileRegex.FindIndex([]byte(p))

	return p[:i[0]-1], p[i[0]:]
}
