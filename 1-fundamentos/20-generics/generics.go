package main

import "fmt"

func SomaInteiro(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}

	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}

	return soma
}

type MyNumber int

type Number interface {
	~int | float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}

	return soma
}

func Compara[T comparable](a, b T) bool {
	return a == b
}
func main() {
	m := map[string]int{"Phellipe": 10000, "Patricia": 20000, "Ping": 1000}
	m2 := map[string]float64{"Phellipe": 10000, "Patricia": 20000, "Ping": 1000}
	m3 := map[string]MyNumber{"Phellipe": 10000, "Patricia": 20000, "Ping": 1000}
	fmt.Printf("SomaInteiro %d\n", SomaInteiro(m))
	fmt.Printf("SomaFloat: %.2f\n", SomaFloat(m2))

	fmt.Printf("Soma m: %d\n", Soma(m))
	fmt.Printf("Soma m2: %.2f\n", Soma(m2))
	fmt.Printf("Soma m3: %d\n", Soma(m3))

	fmt.Println(Compara(10, 10))
}
