package demo2

import (
	"testing"
)

func TestGetArea(t *testing.T) {
	res := GetArea(20, 30)
	if res != 610 {
		t.Error("test failed")
	}
}