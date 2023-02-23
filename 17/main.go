package main

type Number interface {
	int | float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func main() {

	m := map[string]int{"Elvis": 3000, "João": 2000, "Maria": 5000}
	m2 := map[string]float64{"Elvis": 3000.34, "João": 2000.54, "Maria": 5000.67}
	println(Soma(m))
	println(Soma(m2))
}
