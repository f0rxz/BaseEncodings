package base16util

import (
	"bytes"
	"mystdencodings/internal/errchecker"
	"testing"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		input    []byte
		expected string
	}{
		{[]byte{0}, "00"},
		{[]byte{255}, "ff"},
		{[]byte{10, 20, 30}, "0a141e"},
		{[]byte{1, 2, 3, 4, 5}, "0102030405"},
		{[]byte{}, ""},
	}

	for _, test := range tests {
		result := encode(test.input)
		if result != test.expected {
			t.Errorf("Encode(%v) = %s; want %s", test.input, result, test.expected)
		}
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		input          string
		expected       []byte
		expectedErrMsg string
	}{
		{"00", []byte{0}, ""},
		{"ff", []byte{255}, ""},
		{"0a141e", []byte{10, 20, 30}, ""},
		{"0102030405", []byte{1, 2, 3, 4, 5}, ""},
		{"", []byte{}, ""},
		{"0g", nil, "this symbol is not a valid"},
		{"0", nil, "odd length hex string"},
		{"0a141e5", nil, "odd length hex string"},
		{"ffzz", nil, "this symbol is not a valid"},
	}

	for _, test := range tests {
		result, err := decode(test.input)
		if test.expectedErrMsg != "" {
			if err == nil || !errchecker.ContainsError(err.Error(), test.expectedErrMsg) {
				t.Errorf("Decode(%s) error = %v; want %s", test.input, err, test.expectedErrMsg)
			}
		} else {
			if err != nil {
				t.Errorf("Decode(%s) unexpected error: %v", test.input, err)
			}
			if !bytes.Equal(result, test.expected) {
				t.Errorf("Decode(%s) = %v; want %v", test.input, result, test.expected)
			}
		}
	}
}

func TestEncodeDecodeRoundTrip(t *testing.T) {
	tests := [][]byte{
		{0},
		{255},
		{10, 20, 30},
		{1, 2, 3, 4, 5},
		{255, 0, 127, 128, 200},
	}

	for _, test := range tests {
		encoded := encode(test)
		decoded, err := decode(encoded)
		if err != nil {
			t.Errorf("Decode(%s) unexpected error: %v", encoded, err)
		}
		if !bytes.Equal(decoded, test) {
			t.Errorf("Decode(%s) = %v; want %v", encoded, decoded, test)
		}
	}
}
