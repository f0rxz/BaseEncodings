package base32util

import (
	"mystdencodings/internal/errchecker"
	"testing"
)

func TestBase32Encode(t *testing.T) {
	tests := []struct {
		input    []byte
		expected string
	}{
		{[]byte{0}, "AA"},
		{[]byte{255}, "74"},
		{[]byte{10, 20, 30}, "BIKB4"},
		{[]byte{1, 2, 3, 4, 5}, "AEBAGBAF"},
		{[]byte{255, 0, 127, 128, 200}, "74AH7AGI"},
	}

	for _, test := range tests {
		result := encode(test.input)
		if result != test.expected {
			t.Errorf("encode(%v) = %s; want %s", test.input, result, test.expected)
		}
	}
}

func TestBase32Decode(t *testing.T) {
	tests := []struct {
		input          string
		expected       string
		expectedErrMsg string
	}{
		{"AA", string([]byte{0}), ""},
		{"74", string([]byte{255}), ""},
		{"BIKB4", string([]byte{10, 20, 30}), ""},
		{"74AH7AGI", string([]byte{255, 0, 127, 128, 200}), ""},
		{"INVALID###", "", "invalid Base32 data"},
	}

	for _, test := range tests {
		result, err := decode(test.input)
		if test.expectedErrMsg != "" {
			if err == nil || !errchecker.ContainsError(err.Error(), test.expectedErrMsg) {
				t.Errorf("decode(%s) error = %v; want %s", test.input, err, test.expectedErrMsg)
			}
		} else {
			if err != nil {
				t.Errorf("decode(%s) unexpected error: %v", test.input, err)
			}
			if string(result) != test.expected {
				t.Errorf("decode(%s) = %s; want %s", test.input, result, test.expected)
			}
		}
	}
}

func TestBase32EncodeDecodeRoundTrip(t *testing.T) {
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
			t.Errorf("decode(%s) unexpected error: %v", encoded, err)
		}
		if string(decoded) != string(test) {
			t.Errorf("RoundTrip failed for %v: got %v, want %v", test, decoded, test)
		}
	}
}
