package base32util

import (
	"bytes"
	"errors"
	"mystdencodings/internal/delpadd"
	"strings"
)

type Base32Encoder struct{}

func NewEncoder() *Base32Encoder {
	return &Base32Encoder{}
}

func (b *Base32Encoder) Encode(data []byte) string {
	return encode(data)
}

func (b *Base32Encoder) Decode(data string) ([]byte, error) {
	return decode(data)
}

// Base32 alphabet as per RFC 4648
const tobase32 = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"

func buildBase32ReverseMapping() [256]byte {
	var reverseMap [256]byte
	for i := 0; i < 256; i++ {
		reverseMap[i] = 255
	}
	for i := 0; i < len(tobase32); i++ {
		reverseMap[tobase32[i]] = byte(i)
	}
	return reverseMap
}

var Base32Map = buildBase32ReverseMapping()

func encode(data []byte) string {
	var result strings.Builder
	bits := uint(0)
	bitCount := uint(0)

	for _, b := range data {
		bits = (bits << 8) | uint(b)
		bitCount += 8

		for bitCount >= 5 {
			bitCount -= 5
			index := (bits >> bitCount) & 31
			result.WriteByte(tobase32[index])
		}
	}

	if bitCount > 0 {
		bits <<= (5 - bitCount)
		index := bits & 31
		result.WriteByte(tobase32[index])
	}

	padding := (8 - (len(data)*8)%40) % 8
	for padding > 0 {
		result.WriteByte('=')
		padding -= 5
	}

	return result.String()
}

func decode(data string) ([]byte, error) {
	var output bytes.Buffer
	var chr1, chr2, chr3, chr4, chr5 byte
	var enc1, enc2, enc3, enc4, enc5, enc6, enc7, enc8 byte
	i := 0

	data = delpadd.RemovePadding(data)

	for i < len(data) {
		// Read enc1 to enc8 safely
		enc1, enc2, enc3, enc4, enc5, enc6, enc7, enc8 = 32, 32, 32, 32, 32, 32, 32, 32

		if i < len(data) {
			enc1 = Base32Map[data[i]]
			i++
		}
		if i < len(data) {
			enc2 = Base32Map[data[i]]
			i++
		}
		if i < len(data) {
			enc3 = Base32Map[data[i]]
			i++
		}
		if i < len(data) {
			enc4 = Base32Map[data[i]]
			i++
		}
		if i < len(data) {
			enc5 = Base32Map[data[i]]
			i++
		}
		if i < len(data) {
			enc6 = Base32Map[data[i]]
			i++
		}
		if i < len(data) {
			enc7 = Base32Map[data[i]]
			i++
		}
		if i < len(data) {
			enc8 = Base32Map[data[i]]
			i++
		}

		if enc1 == 255 || enc2 == 255 || enc3 == 255 || enc4 == 255 ||
			enc5 == 255 || enc6 == 255 || enc7 == 255 || enc8 == 255 {
			return nil, errors.New("invalid Base32 data")
		}

		chr1 = (enc1 << 3) | (enc2 >> 2)
		output.WriteByte(chr1)

		if enc3 != 32 {
			chr2 = (enc2 << 6) | (enc3 << 1) | (enc4 >> 4)
			output.WriteByte(chr2)
		}

		if enc5 != 32 {
			chr3 = (enc4 << 4) | (enc5 >> 1)
			output.WriteByte(chr3)
		}

		if enc6 != 32 {
			chr4 = (enc5 << 7) | (enc6 << 2) | (enc7 >> 3)
			output.WriteByte(chr4)
		}

		if enc8 != 32 {
			chr5 = (enc7 << 5) | enc8
			output.WriteByte(chr5)
		}
	}

	return output.Bytes(), nil
}
