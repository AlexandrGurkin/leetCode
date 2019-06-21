package atoi

import "testing"

func TestAtoi1(t *testing.T) {
	res := atoi("42")
	if res != 42 {
		t.Error("Expected 42, got ", res)
	}
}

func TestAtoi2(t *testing.T) {
	res := atoi("     -42")
	if res != -42 {
		t.Error("Expected -42, got ", res)
	}
}

func TestAtoi3(t *testing.T) {
	res := atoi("4193 with words")
	if res != 4193 {
		t.Error("Expected 4193, got ", res)
	}
}

func TestAtoi4(t *testing.T) {
	res := atoi("words and 987")
	if res != 0 {
		t.Error("Expected 0, got ", res)
	}
}

func TestAtoi5(t *testing.T) {
	res := atoi("-91283472332")
	if res != -2147483648 {
		t.Error("Expected -2147483648, got ", res)
	}
}

func TestAtoi6(t *testing.T) {
	res := atoi("3.14159")
	if res != 3 {
		t.Error("Expected 3, got ", res)
	}
}

func TestAtoi7(t *testing.T) {
	res := atoi("9223372036854775808")
	if res != 2147483647 {
		t.Error("Expected 2147483647, got ", res)
	}
}

//
func TestAtoi8(t *testing.T) {
	res := atoi("  0000000000012345678")
	if res != 12345678 {
		t.Error("Expected 12345678, got ", res)
	}
}
