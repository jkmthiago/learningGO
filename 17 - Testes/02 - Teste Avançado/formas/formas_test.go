package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T)  {
	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{10, 10}
		expected := float64(100)
		received := ret.Area()

		if received != expected {
			t.Errorf("A área esperada para o quadrado era %f mas recebeu %f", expected, received)
		}
	})

	t.Run("Círculo", func(t *testing.T) {
		circ := Circulo{10}
		expected := float64(100*math.Pi)
		received := circ.Area()

		if received != expected {
			t.Errorf("A área esperada para o circulo era %f mas recebeu %f", expected, received)
		}
	})
}