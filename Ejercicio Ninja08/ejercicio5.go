package main

import (
	"fmt"
	"sort"
)

/*
Comenzando con este código, ordena el []usuario por
	Edad
	Apellido
También ordena el []string “Dichos” para cada usuario
	Imprime todo de una manera agradable

*/

type persona struct {
	Nombre   string
	Apellido string
	Edad     int
	Dichos   []string
}

type porEdad []persona

func (a porEdad) Len() int           { return len(a) }
func (a porEdad) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a porEdad) Less(i, j int) bool { return a[i].Edad < a[j].Edad }

type porApellido []persona

func (a porApellido) Len() int           { return len(a) }
func (a porApellido) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a porApellido) Less(i, j int) bool { return a[i].Apellido < a[j].Apellido }

func ejercicio5() {
	u1 := persona{
		Nombre:   "Eduar",
		Apellido: "Tua",
		Edad:     32,
		Dichos: []string{
			"Cachicamo diciéndole a morrocoy conchudo",
			"La mona, aunque se vista de seda, mona se queda",
			"Poco a poco se anda lejos",
		},
	}

	u2 := persona{
		Nombre:   "Carlos",
		Apellido: "Pérez",
		Edad:     27,
		Dichos: []string{
			"Camarón que se duerme se lo lleva la corriente",
			"A ponerse las alpargatas que lo que viene es joropo",
			"No gastes pólvora en zamuro",
		},
	}

	u3 := persona{
		Nombre:   "Che",
		Apellido: "López",
		Edad:     54,
		Dichos: []string{
			"Ni lava ni presta la batea",
			"Hijo de gato, caza ratón",
			"Más vale pájaro en mano que cien volando",
		},
	}

	usuarios := []persona{u1, u2, u3}
	imprimir(usuarios)

	// tu código va aquí
	// Ordenando los dichos
	for _, v := range usuarios {
		sort.Strings(v.Dichos)
	}
	fmt.Println("Ordenado por Edad -------------------------------------")
	sort.Sort(porEdad(usuarios))
	imprimir(usuarios)
	fmt.Println("Ordenado por Apellido ---------------------------------")
	sort.Sort(porApellido(usuarios))
	imprimir(usuarios)
}

func imprimir(p []persona) {
	for _, v := range p {
		fmt.Println(v.Apellido, v.Nombre)
		fmt.Println(v.Edad)
		for _, vDichos := range v.Dichos {
			fmt.Println("\t", vDichos)
		}
	}
}
