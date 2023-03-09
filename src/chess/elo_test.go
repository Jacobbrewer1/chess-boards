package chess

import (
	"github.com/Jacobbrewer1/chess-boards/src/custom"
	"testing"
)

func TestCalculateNew(t *testing.T) {
	ra := 1200
	rb := 1000
	k := 30
	gotA, gotB := CalculateNew(float64(ra), float64(rb), k, custom.Bool(true))

	//1207.21 Rb = 992.792
	if gotA != float64(1207) {
		t.Errorf("gotA = %f, expected %d", gotA, 1207)
		return
	}

	if gotB != 993 {
		t.Errorf("gotB = %f, expected %d", gotB, 993)
		return
	}
}
