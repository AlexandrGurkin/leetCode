package reverseInteger

import "testing"

func TestReverse(t *testing.T) {
	res := reverse(123)
	if res != 321 {
		t.Error("Expected 321, got ", res)
	}
}

func TestReverseMin(t *testing.T) {
	res := reverse(-123)
	if res != -321 {
		t.Error("Expected 321, got ", res)
	}
}