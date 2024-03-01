package server

type HTTPStatus int

const (
	HTTP_OK HTTPStatus = 200
)

type ResponseHeader struct {
	Status     HTTPStatus
	Additional map[string]string
}

type Response struct {
	Header ResponseHeader
	Body   string
}
