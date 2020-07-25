package myframe

import "net/http"

type (
	Response struct {
		Writer http.ResponseWriter
		Status int
		Size   int64
	}
)

func NewResponse(w http.ResponseWriter) *Response {
	return &Response{Writer: w}
}
