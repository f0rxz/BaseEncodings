package base16util

import "errors"

type Base16Encoder struct{}

func NewEncoder() *Base16Encoder {
	return &Base16Encoder{}
}

func (b *Base16Encoder) Encode(data []byte) string {
	return encode(data)
}

func (b *Base16Encoder) Decode(data string) ([]byte, error) {
	return decode(data)
}

const tobase16 = "0123456789abcdef"

func buildReverseBase64Mapping() [256]byte {
	var reverseMap [256]byte
	for i := 0; i < 256; i++ {
		reverseMap[i] = 255
	}
	for i := 0; i < len(tobase16); i++ {
		reverseMap[tobase16[i]] = byte(i)
	}
	return reverseMap
}

var frombase16 = buildReverseBase64Mapping()

var fromhex = []byte{
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 255, 255, 255, 255, 255, 255,
	255, 10, 11, 12, 13, 14, 15, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 10, 11, 12, 13, 14, 15,
}

func encode(src []byte) string {

	result := make([]byte, len(src)*2)

	for i, key := range src {
		result[2*i] = tobase16[key>>4]
		result[2*i+1] = tobase16[key&15]
	}

	return string(result)
}

func decode(s string) ([]byte, error) {
	result := make([]byte, len(s)/2)

	length := len(s) / 2 * 2
	for i := 0; i < length; i += 2 {
		c := s[i]

		if int(c) >= len(frombase16) {
			return nil, errors.New("this symbol is not a valid")
		}

		h := frombase16[c]

		if h == 255 {
			return nil, errors.New("this symbol is not a valid")
		}

		c = s[i+1]

		if int(c) >= len(frombase16) {
			return nil, errors.New("this symbol is not a valid")
		}

		l := frombase16[c]

		if l == 255 {
			return nil, errors.New("this symbol is not a valid")
		}

		result[i/2] = (h << 4) | l
	}

	if len(s)%2 != 0 {
		return result, errors.New("odd length hex string")
	}
	return result, nil
}
