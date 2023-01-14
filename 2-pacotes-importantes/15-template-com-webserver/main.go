package main

import (
	"html/template"
	"net/http"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		t := template.Must(template.New("template.html").
			ParseFiles("template.html"))
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
