package main

import (
	"fmt"

	"curso-go/matematica"
)

func main() {
	s := matematica.Soma(10, 20)

	fmt.Println("O resultado: ", s)
}
