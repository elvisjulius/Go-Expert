package main

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

func main() {

	m := map[string]int{"Elvis": 3000, "João": 2000, "Maria": 5000}
	m2 := map[string]float64{"Elvis": 3000.34, "João": 2000.54, "Maria": 5000.67}
	println(SomaInteiro(m))
	println(SomaFloat(m2))
}
