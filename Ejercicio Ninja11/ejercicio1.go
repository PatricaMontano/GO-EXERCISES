package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
Comienza con este código https://go.dev/play/p/Bgads38iFgz
En vez de usar el identificador blank (underscore),
asegúrate de que el código esté chequeando y manejando el error.
*/
type persona struct {
	Nombre   string
	Apellido string
	Frases   []string
}

func ejercicio1() {
	p1 := persona{
		Nombre:   "James",
		Apellido: "Bond",
		Frases:   []string{"Shaken, not stirred", "¿Algún último deseo?", "Nunca digas nunca."},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(bs))

}
