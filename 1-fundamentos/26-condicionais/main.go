package main

func main() {
	a := 1
	b := 2
	if a > b {
		println(a)
	} else {
		println(b)
	}

	switch a {
	case 1:
		println("a")
	case 2:
		println("b")
	case 3:
		println("c")
	default:
		println("nenhum")

	}
}
