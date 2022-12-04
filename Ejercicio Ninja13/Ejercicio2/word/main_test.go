package word

import (
	"ejercicio2/quote"
	"fmt"
	"testing"
)

func TestUseCount(t *testing.T) {
	a := UseCount("uno uno dos tres tres tres")
	for k, v := range a {
		switch k {
		case "uno":
			if v != 2 {
				t.Error("Expected", 2, "Got", v)
			}
		case "dos":
			if v != 1 {
				t.Error("Expected", 1, "Got", v)
			}
		case "tres":
			if v != 3 {
				t.Error("Expected", 3, "Got", v)
			}
		}

	}

}

func TestCount(t *testing.T) {
	num := Count(quote.SunAlso)
	if num != 1349 {
		t.Error("Expected", 1349, "Got", num)
	}
}

func ExampleUseCount() {
	fmt.Print(UseCount(quote.SunAlso))
}

func ExampleCount() {
	fmt.Println(Count(quote.SunAlso))
	//Output
	//2
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}

}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}
