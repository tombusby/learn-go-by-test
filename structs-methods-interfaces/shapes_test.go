package shapes

import "testing"

func TestRectandlePerimeter(t *testing.T) {
	t.Run("test perimiter of 10 x 10 rect comes out to 40", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Permiter(rectangle)
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("test perimiter of 20 x 3 rect comes out to 46", func(t *testing.T) {
		rectangle := Rectangle{20.0, 3.0}
		got := Permiter(rectangle)
		want := 46.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

func TestRectandleArea(t *testing.T) {
	t.Run("test area of 10 x 10 rectangle comes out to 100", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Area(rectangle)
		want := 100.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("test area of 20 x 40 rectangle comes out to 800", func(t *testing.T) {
		rectangle := Rectangle{20.0, 40.0}
		got := Area(rectangle)
		want := 800.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}
