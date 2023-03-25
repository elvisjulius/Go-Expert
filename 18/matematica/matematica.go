package matematica

func Soma[T int | float64](a, b T) T {
	return a + b
}

var A int = 10

type Carro struct {
	Marca string
}

func (c Carro) Andar() string {
	return "Carro andando"
}

// letras maiusculas permitem importar uma struct,função ou atributo fora do pacote
