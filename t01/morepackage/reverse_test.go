package morepackage

import "testing"

func TestReverse(t *testing.T) {
	data := []struct {
		input, expected string
	}{
		{"Hello, world!", "!dlrow ,olleH"},
		{"Hello, Go!", "!oG ,olleH"},
	}
	for _, d := range data {
		if got := Reverse(d.input); got != d.expected {
			t.Errorf("Reverse(%q) = %q, want %q", d.input, got, d.expected)
		}
	}
}

func TestConvertUpper(t *testing.T) {
	data := []struct {
		input, expected string
	}{
		{"Hello, world!", "HELLO, WORLD!"},
		{"Hello, Go!", "HELLO, GO!"},
	}
	for _, d := range data {
		if got := ConvertUpper(d.input); got != d.expected {
			t.Errorf("ConvertUpper(%q) = %q, want %q", d.input, got, d.expected)
		}
	}
}
