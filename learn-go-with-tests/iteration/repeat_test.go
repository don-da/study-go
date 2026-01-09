package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Repeat string 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := ("aaaaa")
		assertCorrectRepeat(t, repeated, expected)
	})

	t.Run("Repeat string 10 times", func(t *testing.T) {
		repeated := Repeat("a", 9)
		expected := ("aaaaaaaaa")
		assertCorrectRepeat(t, repeated, expected)
	})
}

func assertCorrectRepeat(t testing.TB, repeated, expected string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("expected %q, got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 3)
	fmt.Println(repeated)
	// Output: aaa
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}
