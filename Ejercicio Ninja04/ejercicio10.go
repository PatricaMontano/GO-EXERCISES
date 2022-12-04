package main

import "fmt"

/*
Usando el código del ejemplo anterior, elimina un registro a tu mapa,
ahora imprime el mapa usando “range”
*/
func ejercicio10() {
	x := map[string][]string{
		`eduar_tua`:    []string{`computadoras`, `montaña`, `playa`},
		`carlos_ramon`: []string{`leer`, `comprar`, `música`},
		`juan_bimba`:   []string{`helado`, `pintar`, `bailar`},
	}
	x[`omar_sanmartin`] = []string{`futbol`, `musica`, `peliculas`}

	delete(x, `eduar_tua`)
	delete(x, `carlos_ramon`)
	for clave, valor := range x {
		fmt.Println("Persona :", clave)
		for i, val := range valor {
			fmt.Println("	Index :", i, "   Valor:", val)
		}
	}
}
