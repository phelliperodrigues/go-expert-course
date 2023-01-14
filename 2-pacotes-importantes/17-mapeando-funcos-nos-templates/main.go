package main

import (
	"html/template"
	"net/http"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {

	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		t := template.New("content.html")
		t.Funcs(template.FuncMap{"ToUpper": ToUpper})

		t = template.Must(t.ParseFiles(templates...))
		err := t.Execute(writer, Cursos{
			{"Go", 40},
			{"Java", 50},
			{"Ruby", 30},
		})
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8082", nil)
}
