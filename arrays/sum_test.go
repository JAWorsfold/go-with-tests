package arrays

import (
	"slices"
	"testing"
)

// go test --cover is very useful being inbuilt
func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	// want := "bob" // compiles

	// DeepEqual is not type safe, see above
	// if !reflect.DeepEqual(got, want) { 

	// can use this as of v1.21
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
