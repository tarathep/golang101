package main

import (
	"testing"
)

func TestInputBase5Hegiht5ShouldBeDisplay(t *testing.T) {
	v := CalTriangle(5, 5)
	if 12.5 != v {
		t.Error("Triangle Area should be '12.5' but have", v)
	}
}
