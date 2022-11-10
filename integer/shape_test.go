package integers

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10, 10)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(10, 10)
	want := 100.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
