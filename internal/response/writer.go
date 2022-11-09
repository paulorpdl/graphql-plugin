package response

import (
	"bytes"
	"net/http"
)

type ResponseWriter struct {
	http.ResponseWriter
	buf *bytes.Buffer
}

func CreateResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		buf:            &bytes.Buffer{},
	}
}

func (r *ResponseWriter) Write(p []byte) (int, error) {
	return r.buf.Write(p)
}

func (r *ResponseWriter) Read() []byte {
	return r.buf.Bytes()
}
