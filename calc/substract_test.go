package calc

import (
	"testing"
)

func TestSubstract(t *testing.T) {
	result := Substract(8, 4)
	if result != 7 {
		t.Error("Add() return Error")
	}
}
