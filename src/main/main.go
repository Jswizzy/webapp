package main

import (
	"bufio"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"viewmodels"
)

func main() {
	templates := populateTemplates()

	http.HandleFunc("/",
		func(writer http.ResponseWriter, request *http.Request) {
			requestedFile := request.URL.Path[1:]
			template := templates.Lookup(requestedFile + ".html")

			var context interface{} = nil
			switch requestedFile {
			case "home":
				context = viewmodels.GetHome()
			case "categories":
				context = viewmodels.GetCategories()
			case "products":
				context = viewmodels.GetProducts()
			case "product":
				context = viewmodels.GetProduct()
			}
			if template != nil {
				template.Execute(writer, context)
			} else {
				writer.WriteHeader(404)
			}
		})

	http.HandleFunc("/img/", serveResource)
	http.HandleFunc("/css/", serveResource)

	http.ListenAndServe(":8000", nil)
}

func serveResource(writer http.ResponseWriter, request *http.Request) {
	path := filepath.Join("public", request.URL.Path)
	var contentType string
	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	} else {
		contentType = "text/plain"
	}

	file, err := os.Open(path)

	if err == nil {
		defer file.Close()
		writer.Header().Add("Content-Type", contentType)
		bufferReader := bufio.NewReader(file)
		bufferReader.WriteTo(writer)
	} else {
		writer.WriteHeader(404)
	}
}

func populateTemplates() *template.Template {
	result := template.New("templates")

	basePath := "templates"
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()

	templatePathsRaw, _ := templateFolder.Readdir(-1)
	templatePaths := new([]string)
	for _, pathInfo := range templatePathsRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths,
				filepath.Join(basePath, pathInfo.Name()))
		}
	}

	result.ParseFiles(*templatePaths...)

	return result
}
