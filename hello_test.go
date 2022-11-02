package Go_learn_with_tests

import "testing"

func TestHello(t *testing.T) {
	got := Hello("World")
	want := "Hello World"
	if got != want {
		t.Errorf("We got a %q, but we want a %q", got, want)
	}
}
