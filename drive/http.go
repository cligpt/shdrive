package drive

import (
	"io"
	"net/http"
)

func (d *drive) getRoot(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "root\n")
}

func (d *drive) getHello(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "hello\n")
}
