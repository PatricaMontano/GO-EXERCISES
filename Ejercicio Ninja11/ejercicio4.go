package main

import (
	"fmt"
	"log"
)

/*
Comenzando con este código https://go.dev/play/p/jjfF_jsGplU
Usa el struct raiz.Error como valor de tipo error. Si quieres usa estos valores para tu
lat "50.2289 N"
long "99.4656 W"
*/
type raizError struct {
	lat  string
	long string
	err  error
}

func (re raizError) Error() string {
	return fmt.Sprintf("Error matemático: %v %v %v", re.lat, re.long, re.err)
}

// Ejercicio4 es na
func ejercicio4() {
	_, err := raiz(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func raiz(f float64) (float64, error) {
	if f < 0 {
		// Escribe tu código aquí
		er := fmt.Errorf("El numero %v tiene que ser mayor a 0", f)
		return 0, raizError{"50.2289 N", "99.4656 W", er}
	}
	return 42, nil
}
