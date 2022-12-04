package starting_code

/*
Comienza con este código. Obtenga el código listo para hacer BET en él
(agregue benchmarks, examples, tests). Ejecute lo siguiente en este orden:
tests
benchmarks
coverage
coverage mostrado en el navegador
examples mostrados en documentación en el navegador web

*/
import (
	"ejercicio1/dog"
	"fmt"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}
