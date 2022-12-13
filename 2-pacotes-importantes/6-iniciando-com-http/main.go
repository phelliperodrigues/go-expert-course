package main

import "net/http"

func main() {
	http.HandleFunc("/", BuscarCEP)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func BuscarCEP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello World!"))
	if err != nil {
		panic(err)
	}
}
