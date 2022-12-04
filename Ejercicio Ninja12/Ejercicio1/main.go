package main

/*
Crea un paquete perro. El paquete perro debería tener una función exportada
“Edad” el cuál toma años humano y los retorna como años de perro
(1 año humano = 7 años de perro). Documenta tu código con comentarios.
Usa este código en func main.
Corre el programa y asegúrate de que funciona.

*/
import (
	"fmt"
	"migo/perro"
)

func main() {
	//ejercicio1()
	a := 7
	fmt.Println("La edad de mi perro es de", perro.Edad(a))
}
