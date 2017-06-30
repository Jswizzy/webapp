package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "text/html")
		templates := template.New("template")
		templates.New("test").Parse(doc)
		templates.New("header").Parse(header)
		templates.New("footer").Parse(footer)
		context := Context{
			[3]string{"Lemon", "Orange", "Apple"},
			"the title",
		}
		templates.Lookup("test").Execute(writer, context)

	})

	http.ListenAndServe(":8000", nil)
	//http.ListenAndServe(":8000", http.FileServer(http.Dir("public")))
}

const doc = `
{{template "header" .Title}}
	<body>
		<h1>List of Fruit</h1>
		<ul>
			{{range .Fruit}}
				<li>{{.}}</li>
			{{end}}
		</ul>
	</body>
{{template "footer"}}
`
const header = `
<!DOCTYPE html>
<html>
	<head><title>{{.}}</title></head>
`

const footer = `
</html>
`

type Context struct {
	Fruit [3]string
	Title string
}
