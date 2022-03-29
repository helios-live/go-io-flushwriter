package flusher // import go.ideatocode.tech/io/flushwriter

import (
	"io"
	"net/http"
)

// New returns a new instance of Writer wrapped around w
func New(w io.Writer) Writer {
	return Writer{w: w}
}

// Writer flushes after every call to Writer
type Writer struct {
	w io.Writer
}

func (fw Writer) Write(p []byte) (n int, err error) {
	n, err = fw.w.Write(p)
	fw.w.(http.Flusher).Flush()
	return n, err
}

// Flush writes the pending bytes
func (fw Writer) Flush() {
	fw.w.(http.Flusher).Flush()
}
