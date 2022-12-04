package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
Comenzando con este código, codifica a JSON el []usuario enviando el resultado a Stdout.
Pista: Necesitarás usar json.NewEncoder(os.Stdout).encode(v interface{})
*/

type usuario3 struct {
	Nombre   string
	Apellido string
	Edad     int
	Dichos   []string
}

func ejercicio3() {
	u1 := usuario3{
		Nombre:   "Eduar",
		Apellido: "Tua",
		Edad:     32,
		Dichos: []string{
			"Cachicamo diciéndole a morrocoy conchudo",
			"La mona, aunque se vista de seda, mona se queda",
			"Poco a poco se anda lejos",
		},
	}

	u2 := usuario3{
		Nombre:   "Carlos",
		Apellido: "Pérez",
		Edad:     27,
		Dichos: []string{
			"Camarón que se duerme se lo lleva la corriente",
			"A ponerse las alpargatas que lo que viene es joropo",
			"No gastes pólvora en zamuro",
		},
	}

	u3 := usuario3{
		Nombre:   "Che",
		Apellido: "López",
		Edad:     54,
		Dichos: []string{
			"Ni lava ni presta la batea",
			"Hijo de gato, caza ratón",
			"Más vale pájaro en mano que cien volando",
		},
	}

	usuarios := []usuario3{u1, u2, u3}

	fmt.Println(usuarios)

	// Tu código va aquí
	err := json.NewEncoder(os.Stdout).Encode(usuarios)
	if err != nil {
		fmt.Println("Hicimos algo mal, este es el error:", err)
	}
}
