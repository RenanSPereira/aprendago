package main

import (
	"fmt"
)

// Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".

// 42
// "James Bond"
// true
// Agora demonstre os valores nestas variáveis, com:

// Uma única declaração print
// Múltiplas declarações print

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Print(x, y, z)

	fmt.Print(x)
	fmt.Print(y)
	fmt.Print(z)
}
