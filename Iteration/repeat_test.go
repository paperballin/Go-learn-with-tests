package Iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Repeat a character", func(t *testing.T) {
		got := Repeat("a", 5)
		expect := "aaaaa"
		assertMessage(t, got, expect)
	})
}

func BenchmarkRepeat(b *testing.B) {
	// Execute b.N times
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}

func assertMessage(t testing.TB, got string, expect string) {
	if got != expect {
		t.Errorf("Got: %q, Expect: %q", got, expect)
	}
}
