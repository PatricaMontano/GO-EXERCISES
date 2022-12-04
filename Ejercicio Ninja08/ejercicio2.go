package main

import (
	"encoding/json"
	"fmt"
)

/*
Comenzando con este código, unmarshal el JSON a un estructura de datos de Go.
Pista: https://mholt.github.io/json-to-go/
*/

type autor struct {
	Nombre   string   `json:"Nombre"`
	Apellido string   `json:"Apellido"`
	Edad     int      `json:"Edad"`
	Dichos   []string `json:"Dichos"`
}

func ejercicio2() {
	s := `[{"Nombre":"Eduar","Apellido":"Tua","Edad":32,"Dichos":["Cachicamo diciéndole a morrocoy conchudo","La mona, aunque se vista de seda, mona se queda","Poco a poco se anda lejos"]},{"Nombre":"Carlos","Apellido":"Pérez","Edad":27,"Dichos":["Camarón que se duerme se lo lleva la corriente","A ponerse las alpargatas que lo que viene es joropo","No gastes pólvora en zamuro"]},{"Nombre":"M","Apellido":"Hmmmm","Edad":54,"Dichos":["Ni lava ni presta la batea","Hijo de gato, caza ratón","Más vale pájaro en mano que cien volando"]}]`
	fmt.Println(s)
	//tu código va aquí
	fmt.Println("-------------------------")
	bs := []byte(s)
	var autores []autor
	err := json.Unmarshal(bs, &autores)
	if err != nil {
		fmt.Println(err)
	}

	for i, v := range autores {
		fmt.Println("\nAutor ", i+1)
		fmt.Println(v.Nombre, v.Apellido, v.Edad)
		for _, v := range v.Dichos {
			fmt.Println("\t", v)
		}
	}
}
