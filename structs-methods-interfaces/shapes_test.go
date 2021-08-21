package shapes

import "testing"

func TestRectandlePerimeter(t *testing.T) {
	t.Run("test perimeter of 10 x 10 rect comes out to 40", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("test perimeter of 20 x 3 rect comes out to 46", func(t *testing.T) {
		rectangle := Rectangle{20.0, 3.0}
		got := rectangle.Perimeter()
		want := 46.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

func TestRectandleArea(t *testing.T) {
	t.Run("test area of 10 x 10 rectangle comes out to 100", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area()
		want := 100.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("test area of 20 x 40 rectangle comes out to 800", func(t *testing.T) {
		rectangle := Rectangle{20.0, 40.0}
		got := rectangle.Area()
		want := 800.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

func TestCirclePerimeter(t *testing.T) {
	t.Run("test correct perimeter of r=10 ciricle", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Perimeter()
		want := 62.83185307179586

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}

func TestCircleArea(t *testing.T) {
	t.Run("test correct area of r=10 ciricle", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}
