package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	got := Permiter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Run("test that 10 x 10 comes out to 100", func(t *testing.T) {
		got := Area(10.0, 10.0)
		want := 100.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("test that 20 x 40 comes out to 800", func(t *testing.T) {
		got := Area(20.0, 40.0)
		want := 800.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}
