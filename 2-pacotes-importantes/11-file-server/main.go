package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/blog", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Blog"))
	})
	log.Fatal(http.ListenAndServe(":8080", mux))

}
