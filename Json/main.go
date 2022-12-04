package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	Nombre   string
	Apellido string
	edad     int
}

func main() {
	p1 := persona{
		Nombre:   "James",
		Apellido: "Bond",
		edad:     32,
	}

	p2 := persona{
		Nombre:   "Miss",
		Apellido: "Moneypenny",
		edad:     27,
	}

	personas := []persona{p1, p2}

	fmt.Println(personas)

	bs, err := json.Marshal(personas)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}