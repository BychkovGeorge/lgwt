package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Summ of all numbers in []int{1, 2, 3, 4, 5}", func(t *testing.T) {
		sum := Sum([]int{1, 2, 3, 4, 5})
		expected := 15
		if sum != expected {
			t.Errorf("Sum was incorrect, got: %d, want: %d.", sum, expected)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Summ slices []int{1, 2, 3} and []int{2, 3, 4}", func(t *testing.T) {
		sum := SumAll([]int{1, 2, 3}, []int{2, 3, 4})
		expected := []int{6, 9}
		if !slices.Equal(sum, expected) {
			t.Errorf("Got %v, expected %v", sum, expected)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("Summ of tails of slices []int{1, 2, 3} and []int{2, 3, 4}", func(t *testing.T) {
		sum := SumAllTails([]int{1, 2, 3}, []int{2, 3, 4})
		expected := []int{5, 7}
		if !slices.Equal(sum, expected) {
			t.Errorf("Got %v, expected %v", sum, expected)
		}
	})
}
