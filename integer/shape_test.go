package integers

import "testing"

type Shape interface {
	Area() float64
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10, 10}

	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

}

func TestAreaNew(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{12, 6}, want: 72.0},
		{shape: Circle{10}, want: 314.1592653589793},
		{shape: Triangle{10, 10}, want: 50.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {

			t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
		}
	}

}
