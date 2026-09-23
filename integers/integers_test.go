package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("Test adding two integers", func(t *testing.T) {
		expected := 5
		got := Add(3, 2)
		if expected != got {
			t.Errorf("Add(3, 2) expected to be %d, but got %d", expected, got)
		}
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
