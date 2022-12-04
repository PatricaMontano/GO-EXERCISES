package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	n := Years(7)
	if n != 49 {
		t.Error("Expected", 49, "Got", n)
	}
}

func TestYearsTwo(t *testing.T) {
	n := YearsTwo(7)
	if n != 49 {
		t.Error("Expected", 49, "Got", n)
	}
}

func ExampleYears() {
	fmt.Println(Years(7))
	//Output
	//49
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(7))
	//Output
	//49
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}
