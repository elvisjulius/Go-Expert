package main

import (
	"fmt"
)

type Cliente struct {
	nome string
}

func (c *Cliente) andou() {
	c.nome = "Elvis Julius"
	fmt.Printf("O cliente %v andou\n", c.nome)
}
func main() {
	elvis := Cliente{
		nome: "Elvis",
	}
	elvis.andou()
	fmt.Printf("o valor da struck com nome %v", elvis.nome)
}
