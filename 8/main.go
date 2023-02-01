package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := sum(5, 10)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(valor)
	}

}

func sum(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("A Soma é maior que 50")
	}
	return a + b, nil
}
