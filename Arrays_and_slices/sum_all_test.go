package Arrays_and_slices

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{5, 6, 7, 8, 9})
	want := []int{3, 35}
	t.Run("To sum up all the slices", func(t *testing.T) {
		assertMessageSumAll(t, got, want)
	})
}

func assertMessageSumAll(t testing.TB, got, want []int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got:%v, Want:%v", got, want)
	}
}
