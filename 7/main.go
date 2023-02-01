package main

import "fmt"

func main() {
	salarios := map[string]int{"Elvis": 1000, "Joao": 1500, "Wesley": 2000}
	delete(salarios, "Elvis")
	salarios["Elvis"] = 3000

	// sal := make(map[string]int) inicializar com make
	// sal1 := map[string]int{} iniciliazar diretamente com map
	// sal1["Elvis"] = 5000 colocar um novo valor
	// sal["Marieli"] = 3000

	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salário é de %d\n", salario)
	}

}
