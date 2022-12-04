package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type prueba struct {
		datos     []int
		respuesta float64
	}
	pruebas := []prueba{
		prueba{[]int{8, 8, 8, 10, 10, 10}, 9},
		prueba{[]int{10, 5, 10, 5, 10, 5}, 7.5},
		prueba{[]int{0, 7, 5, 8, 4, 6, 8, 9, 9, 100}, 7},
		prueba{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5},
	}

	for _, v := range pruebas {
		a := CenteredAvg(v.datos)
		if a != v.respuesta {
			t.Error("Expected", v.respuesta, "Got", a)
		}

	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{8, 10, 8, 10}))
	//Output
	//9
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10000000})
	}
}
