package main

func main() {
	// Memória -> Endereço -> Valor
	a := 10

	var ponteiro *int = &a
	*ponteiro = 20
	b := &a
	println(a)
	//imprime o valor
	println(*b)
	//imprime o endereço na me memória
	println(&a)
}
