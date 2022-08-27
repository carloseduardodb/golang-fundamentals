package shapes

import (
	"math"
	"testing"
)

func TestAria(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		rect := Rectangle{10, 12}
		expectAria := 120.0
		actualAria := rect.Aria()
		if actualAria != expectAria {
			t.Fatalf("Expected aria %f, got %f", expectAria, actualAria)
		}
	})

	t.Run("Square", func(t *testing.T) {
		square := Square{10}
		expectAria := float64(math.Pi * 100)
		actualAria := square.Aria()
		if actualAria != expectAria {
			t.Fatalf("Expected aria %f, got %f", expectAria, actualAria)
		}
	})
}
