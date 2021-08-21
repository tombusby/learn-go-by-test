package shapes

import "testing"

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.00},
		{Rectangle{10, 60}, 600},
		{Circle{10}, 314.1592653589793},
		{Circle{24}, 1809.5573684677208},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10.0, 10.0}, 40.0},
		{Rectangle{20.0, 3.0}, 46.0},
		{Circle{10}, 62.83185307179586},
		{Circle{54}, 339.29200658769764},
	}

	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
