package shapes

import "testing"

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle1 := Rectangle{12, 6}
		checkArea(t, rectangle1, 72.00)

		rectangle2 := Rectangle{10, 60}
		checkArea(t, rectangle2, 600)
	})

	t.Run("circles", func(t *testing.T) {
		circle1 := Circle{10}
		checkArea(t, circle1, 314.1592653589793)

		circle2 := Circle{24}
		checkArea(t, circle2, 1809.5573684677208)
	})
}

func TestPerimeter(t *testing.T) {

	checkPerimeter := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle1 := Rectangle{10.0, 10.0}
		checkPerimeter(t, rectangle1, 40.0)

		rectangle2 := Rectangle{20.0, 3.0}
		checkPerimeter(t, rectangle2, 46.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle1 := Circle{10}
		checkPerimeter(t, circle1, 62.83185307179586)

		circle2 := Circle{54}
		checkPerimeter(t, circle2, 339.29200658769764)
	})
}
