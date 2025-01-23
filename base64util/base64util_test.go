package base64util

import (
	"bytes"
	"encoding/base64"
	"mystdencodings/internal/errchecker"
	"testing"
)

func TestBase64Encode(t *testing.T) {
	tests := []struct {
		input    []byte
		expected string
	}{
		{[]byte{0}, "AA=="},
		{[]byte{255}, "/w=="},
		{[]byte{10, 20, 30}, "ChQe"},
		{[]byte{1, 2, 3, 4, 5}, "AQIDBAU="},
		{[]byte{255, 0, 127, 128, 200}, "/wB/gMg="},
	}

	for _, test := range tests {
		result := base64.StdEncoding.EncodeToString(test.input)
		if result != test.expected {
			t.Errorf("encode(%v) = %s; want %s", test.input, result, test.expected)
		}
	}
}

func TestBase64Decode(t *testing.T) {
	tests := []struct {
		input          string
		expected       []byte
		expectedErrMsg string
	}{
		{"AA==", []byte{0}, ""},
		{"/w==", []byte{255}, ""},
		{"ChQe", []byte{10, 20, 30}, ""},
		{"AQIDBAU=", []byte{1, 2, 3, 4, 5}, ""},
		{"/wB/gMg=", []byte{255, 0, 127, 128, 200}, ""},
		{"INVALID###", nil, "illegal base64 data at input byte 7"},
	}

	for _, test := range tests {
		result, err := base64.StdEncoding.DecodeString(test.input)
		if test.expectedErrMsg != "" {
			if err == nil || !errchecker.ContainsError(err.Error(), test.expectedErrMsg) {
				t.Errorf("decode(%s) error = %v; want %s", test.input, err, test.expectedErrMsg)
			}
		} else {
			if err != nil {
				t.Errorf("decode(%s) unexpected error: %v", test.input, err)
			}
			if !bytes.Equal(result, test.expected) {
				t.Errorf("decode(%s) = %v; want %v", test.input, result, test.expected)
			}
		}
	}
}

func TestBase64EncodeDecodeRoundTrip(t *testing.T) {
	tests := [][]byte{
		{0},
		{255},
		{10, 20, 30},
		{1, 2, 3, 4, 5},
		{255, 0, 127, 128, 200},
	}

	for _, test := range tests {
		encoded := base64.StdEncoding.EncodeToString(test)
		decoded, err := base64.StdEncoding.DecodeString(encoded)
		if err != nil {
			t.Errorf("decode(%s) unexpected error: %v", encoded, err)
		}
		if !bytes.Equal(decoded, test) {
			t.Errorf("RoundTrip failed for %v: got %v, want %v", test, decoded, test)
		}
	}
}
