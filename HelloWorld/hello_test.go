package HelloWorld

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say hello to everyone", func(t *testing.T) {
		got := Hello("", "Chinese")
		want := "你好, 世界"
		assertMessage(t, got, want)
	})
}

func assertMessage(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("We got a %q, but we want a %q", got, want)
	}
}
