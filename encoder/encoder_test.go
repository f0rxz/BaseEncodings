package encoder

import (
	"bytes"
	"testing"
)

func TestEncoder_EncodeDecode(t *testing.T) {
	tests := []struct {
		name           string
		input          []byte
		expectedBase64 string
		expectedBase32 string
		expectedBase16 string
	}{
		{
			name:           "Test1",
			input:          []byte{255, 0, 127, 128, 200},
			expectedBase64: "/wB/gMg=",
			expectedBase32: "74AH7AGI",
			expectedBase16: "ff007f80c8",
		},
		{
			name:           "Test2",
			input:          []byte{0},
			expectedBase64: "AA==",
			expectedBase32: "AA",
			expectedBase16: "00",
		},
		{
			name:           "Test3",
			input:          []byte{1, 2, 3, 4, 5},
			expectedBase64: "AQIDBAU=",
			expectedBase32: "AEBAGBAF",
			expectedBase16: "0102030405",
		},
		{
			name:           "Test4",
			input:          []byte{10, 20, 30},
			expectedBase64: "ChQe",
			expectedBase32: "BIKB4",
			expectedBase16: "0a141e",
		},
		{
			name:           "Test5",
			input:          []byte{},
			expectedBase64: "",
			expectedBase32: "",
			expectedBase16: "",
		},
	}

	encoder := NewEncoder()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			base64Result := encoder.Base64.Encode(test.input)
			if base64Result != test.expectedBase64 {
				t.Errorf("Base64 encode failed for %v: got %s, want %s", test.input, base64Result, test.expectedBase64)
			}
			base64Decoded, err := encoder.Base64.Decode(base64Result)
			if err != nil || !bytes.Equal(base64Decoded, test.input) {
				t.Errorf("Base64 decode failed for %s: got %v, want %v", base64Result, base64Decoded, test.input)
			}

			base32Result := encoder.Base32.Encode(test.input)
			if base32Result != test.expectedBase32 {
				t.Errorf("Base32 encode failed for %v: got %s, want %s", test.input, base32Result, test.expectedBase32)
			}
			base32Decoded, err := encoder.Base32.Decode(base32Result)
			if err != nil || !bytes.Equal(base32Decoded, test.input) {
				t.Errorf("Base32 decode failed for %s: got %v, want %v", base32Result, base32Decoded, test.input)
			}

			base16Result := encoder.Base16.Encode(test.input)
			if base16Result != test.expectedBase16 {
				t.Errorf("Base16 encode failed for %v: got %s, want %s", test.input, base16Result, test.expectedBase16)
			}
			base16Decoded, err := encoder.Base16.Decode(base16Result)
			if err != nil || !bytes.Equal(base16Decoded, test.input) {
				t.Errorf("Base16 decode failed for %s: got %v, want %v", base16Result, base16Decoded, test.input)
			}
		})
	}
}
