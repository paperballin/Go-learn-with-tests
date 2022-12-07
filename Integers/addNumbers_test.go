package Integers

import "testing"

func TestAddNumbers(t *testing.T) {
	sum := Add(2, 2)
	expect := 4
	assertMessage(t, expect, sum)

}

func assertMessage(t testing.TB, expect int, sum int) {
	if sum != expect {
		t.Errorf("Expect: %d, Got: %d", expect, sum)
	}
}
