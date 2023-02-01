package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(1, 3, 45, 6, 323, 6262, 656, 565, 656, 456, 656, 656))

}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
