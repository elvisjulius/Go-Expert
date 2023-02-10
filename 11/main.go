package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado e seu status é %t\n", c.Nome, c.Ativo)
}

func main() {
	wesley := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
	}

	fmt.Println(wesley.Ativo)
	wesley.Desativar()
	fmt.Println(wesley.Ativo)
	wesley.Endereco.Cidade = "São Paulo"
}
