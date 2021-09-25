package calc

import (
	"testing"
)

func TestSubstract(t *testing.T) {
	result := Substract(8, 4)
	if result != 4 {
		t.Error("Substract() return Error")
	}
}
