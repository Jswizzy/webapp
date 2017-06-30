package main

import (
	"net/http"
	"path/filepath"
	"strings"
	"os"
	"bufio"
)

func main() {
	http.Handle("/", new(MyHandler))

	http.ListenAndServe(":8000", nil)
}

type MyHandler struct {
	http.Handler
}

func (myHandler *MyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	path := filepath.Join("public",  request.URL.Path)
	file, err := os.Open(path)

	if err == nil {
		bufferedReader := bufio.NewReader(file)
		var contentType string

		if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(path, ".mp4") {
			contentType = "video/mp4"
		} else {
			contentType = "text/plain"
		}

		writer.Header().Add("Content-Type", contentType)
		bufferedReader.WriteTo(writer)
	} else {
		writer.WriteHeader(404)
		writer.Write([]byte("404 - " + http.StatusText(404)))
	}
}