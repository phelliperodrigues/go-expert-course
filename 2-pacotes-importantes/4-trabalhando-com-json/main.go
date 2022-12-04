package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int `json:"numero"`
	Saldo  int `json:"saldo"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}

	println(string(res))

	encoder := json.NewEncoder(os.Stdout)
	err = encoder.Encode(conta)
	if err != nil {
		println(err)
	}

	var contaX Conta
	jsonString := []byte(`{"numero":2,"saldo":200}`)
	err = json.Unmarshal(jsonString, &contaX)
	if err != nil {
		println(err)
	}
	fmt.Printf("%#v\n", contaX)

}
