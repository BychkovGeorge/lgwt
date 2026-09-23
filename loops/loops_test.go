package loops

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Repeat symbol 'a' five times", func(t *testing.T) {
		res := Repeat("a", 5)
		expected := "aaaaa"
		if res != expected {
			t.Errorf("Expected to get %q as a result, but got %q", expected, res)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	sum := Repeat("a", 5)
	fmt.Println(sum)
	// Result: "aaaaa"
}
