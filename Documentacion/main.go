// Package mymath provides math solutions
package mymath

//Sum suma es un funcion que suma valores ilimitados
func Sum(xi ...int) int {
	sum := 0
	for v := range xi {
		sum += v
	}
	return sum
}
