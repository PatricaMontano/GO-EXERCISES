package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
Comienza con este código https://go.dev/play/p/kbeQfbe82Nl
Crea un mensaje de error personalizado usando “fmt.Errorf”.
*/
func ejercicio2() {
	p1 := persona{
		Nombre:   "James",
		Apellido: "Bond",
		Frases:   []string{"Shaken, not stirred", "¿Algún último deseo?", "Nunca digas nunca."},
	}

	bs, err := aJSON(p1)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(bs))

}

// aJSON también necesita retorna un error
func aJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("Hubo un error en aJSON: %v", err)
	}
	return bs, nil
}
