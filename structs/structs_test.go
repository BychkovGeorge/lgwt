package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("Calculate rectangle perimeter", func(t *testing.T) {
		rectangle := Rectangle{10, 20}
		got := rectangle.Perimeter()
		want := 60.00
		if got != want {
			t.Errorf("Perimeter is %.2f, but has to be %.2f", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{base: 10, height: 13}, 65},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
