// Package perro nos permite entender mejor los perros
package perro

import "fmt"

// Edad nos devuelve la edad de un perro en años perro
func Edad(e int) int {
	return e * 7
}

func main() {
	fmt.Println("La edad del humano es ", 7, " y la del perro es ", Edad(7))
}
