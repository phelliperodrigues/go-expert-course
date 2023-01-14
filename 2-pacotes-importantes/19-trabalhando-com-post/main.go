package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	jsonVar := bytes.NewBuffer([]byte(`{"title": "phellipe"}`))
	resp, err := c.Post("https://dummyjson.com/products/add", "application/json", jsonVar)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil)

}
