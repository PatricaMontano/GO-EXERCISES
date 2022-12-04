package main

import (
	"ejercicio2/quote"
	"ejercicio2/word"
	"fmt"
)

/*
Comienza con este código. https://github.com/GoesToEleven/go-programming/tree/master/code_samples/010-ninja-level-thirteen/02/01-code-starting/word
Obtén el código listo para BET
(agrega benchmarks, examples, tests) sin embargo no escribas un ejemplo para la
función que retorna un mapa; solamente escribe una prueba para ella como un reto extra.
Agrégale documentación al código. Sigue el siguiente orden:
tests
benchmarks
coverage
coverage shown in web browser
examples shown in documentation in a web browser
*/
func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
