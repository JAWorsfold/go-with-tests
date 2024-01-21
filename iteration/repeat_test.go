package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	repeat := Repeat("j", 5)
	fmt.Println(repeat)
	// Output: jjjjj
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
