package saludar

import "fmt"

// Greet nos permite retornar un saludo
func Greet(s string) string {
	return fmt.Sprint("Hello my dear, ", s)
}
