package arrays_slices

import (
	"reflect"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		assertSum(t, got, want, numbers)
	})
}

func assertSum(t testing.TB, got, want int, given []int) {
	if got != want {
		t.Errorf("got %d, want %d, given %d", got, want, given)
	}
}

func TestSumAll(t *testing.T) {
	t.Run("return a slice with the sum of n slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{3, 4})
		want := []int{3, 7}

		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{3, 4})
	want := []int{2, 4}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
