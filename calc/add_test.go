package calc

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(3, 4)
	if result != 7 {
		t.Error("Add() return Error")
	}
}
