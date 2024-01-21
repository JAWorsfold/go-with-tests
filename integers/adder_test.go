package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAdder(t *testing.T) {
	t.Run("add 2 and 2", func(t *testing.T) {
		received := Add(2, 2)
		expected := 4
		assertCorrectMessage(t, received, expected)
	})

	t.Run("add -1 and 3", func(t *testing.T) {
		received := Add(-1, 3)
		expected := 2
		assertCorrectMessage(t, received, expected)
	})

	t.Run("add 0 and 33", func(t *testing.T) {
		received := Add(0, 33)
		expected := 33
		assertCorrectMessage(t, received, expected)
	})
}

func assertCorrectMessage(t testing.TB, received, expected int) {
	t.Helper()
	if received != expected {
		t.Errorf("expected '%d' but got '%d'", expected, received)
	}
}
