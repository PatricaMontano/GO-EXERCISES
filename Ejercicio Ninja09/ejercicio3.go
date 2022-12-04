package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
Usando gorutinas, crea un programa incremento

	Haz que una variable guarde el valor del incremento
	Lanza varias gorutinas
		cada gorutina deberá
			Leer el valor del incremento
			Almacenarlo en una nueva variable
			Ceder el procesador con runtime.Gosched()
			Incrementa la nueva variable
			Escribe el valor en la nueva variable de vuelta a la variable incremento

Usa waitgroups para esperar que todas las gorutinas finalicen
Lo anterior generará una race condition.
Comprueba que es una race condition usando el flag -race
Si necesitas ayuda, aquí tienes una pista: https://play.golang.org/p/a-tdD-7lTId
*/
func ejercicio3() {
	fmt.Println("Número de Goroutinas:", runtime.NumGoroutine())
	contador := 0
	const incremento = 50

	var wg sync.WaitGroup
	wg.Add(incremento)

	for i := 0; i < incremento; i++ {
		go func() {
			defer wg.Done()
			aux := contador
			runtime.Gosched()
			aux++
			contador = aux
		}()
		fmt.Println("Número de Goroutinas:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("El contador final es", contador)

}
