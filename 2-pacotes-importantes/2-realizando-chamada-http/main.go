package main

import (
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://www.google.com.br")
	if err != nil {
		panic(err)
	}
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	println(string(res))

	err = req.Body.Close()
	if err != nil {
		return
	}
}
