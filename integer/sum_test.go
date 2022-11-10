package integers

import (
	"testing"
)

func TestSum(t *testing.T) {

	numbers := []int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	assertCorrectMessage(t, got, want, numbers)
}

func assertCorrectMessage(t *testing.T, got int, want int, numbers []int) {
	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
