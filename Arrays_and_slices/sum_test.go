package Arrays_and_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers1 := []int{1, 2, 3, 4, 5}
	numbers2 := []int{1, 2, 3}
	want1 := 15
	want2 := 6
	t.Run("This is a test for two numbers", func(t *testing.T) {
		got1 := Sum(numbers1)
		got2 := Sum(numbers2)
		assertMessageSum(t, got1, want1)
		assertMessageSum(t, got2, want2)
	})
}

func assertMessageSum(t testing.TB, got, want int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got: %d, Want: %v", got, want)
	}
}
