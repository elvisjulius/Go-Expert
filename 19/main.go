package main

func main() {

	//For comum
	for i := 0; i < 5; i++ {
		println(i)
	}

	numeros := []string{"um", "dois", "três"}

	// utilizando key e value
	for k, v := range numeros {
		println(k, v)
	}

	//utilizando blank identifier para receber apenas o value
	for _, v := range numeros {
		println(v)
	}

	//utilizando blank identifier para receber apenas o key
	for k, _ := range numeros {
		println(k)
	}

	//simplificando
	j := 0
	for j < 5 {
		println(j)
		j++
	}

	//looping infinito da pra usar pra consumir uma fila por exemplo
	for {
		println("Hello, World")
	}

}
