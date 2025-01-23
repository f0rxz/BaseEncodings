package base64util

import (
	"bytes"
	"errors"
	"mystdencodings/internal/delpadd"
)

/*
function base64_encode( data ) {	// Encodes data with MIME base64
	//
	// +   original by: Tyler Akins (http://rumkin.com)
	// +   improved by: Bayron Guevara

	var b64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=";
	var o1, o2, o3, h1, h2, h3, h4, bits, i=0, enc='';

	do { // pack three octets into four hexets
		o1 = data.charCodeAt(i++); //returns zero if index out of range thats our problem
		o2 = data.charCodeAt(i++);
		o3 = data.charCodeAt(i++);

		bits = o1<<16 | o2<<8 | o3;

		h1 = bits>>18 & 0x3f; //0x3f == 3*16+15==63
		h2 = bits>>12 & 0x3f;
		h3 = bits>>6 & 0x3f;
		h4 = bits & 0x3f;

		// use hexets to index into b64, and append result to encoded string
		enc += b64.charAt(h1) + b64.charAt(h2) + b64.charAt(h3) + b64.charAt(h4);
	} while (i < data.length);

	switch( data.length % 3 ){
		case 1:
			enc = enc.slice(0, -2) + '==';
		break;
		case 2:
			enc = enc.slice(0, -1) + '=';
		break;
	}

	return enc;
}
*/

type Base64Encoder struct{}

func NewEncoder() *Base64Encoder {
	return &Base64Encoder{}
}

func (b *Base64Encoder) Encode(data []byte) string {
	return encode(data)
}

func (b *Base64Encoder) Decode(data string) ([]byte, error) {
	return decode(data)
}

//this is encoded map for hex generate encode map for base 64

const tobase64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func buildReverseBase64Mapping() [256]byte {
	var reverseMap [256]byte
	for i := 0; i < 256; i++ {
		reverseMap[i] = 255
	}
	for i := 0; i < len(tobase64); i++ {
		reverseMap[tobase64[i]] = byte(i)
	}
	return reverseMap
}

var base64Map = buildReverseBase64Mapping()

// func encodeFromJS(data []byte) string {
// 	var o1, o2, o3, h1, h2, h3, h4, bits int
// 	var enc string

// 	i := 0
// 	for i < len(data) {
// 		o1 = int(data[i])
// 		i++
// 		o2 = 0
// 		o3 = 0
// 		if i < len(data) {
// 			o2 = int(data[i])
// 			i++
// 		}
// 		if i < len(data) {
// 			o3 = int(data[i])
// 			i++
// 		}

// 		bits = (o1 << 16) | (o2 << 8) | o3

// 		h1 = (bits >> 18) & 0x3F
// 		h2 = (bits >> 12) & 0x3F
// 		h3 = (bits >> 6) & 0x3F
// 		h4 = bits & 0x3F

// 		enc += string(tobase64[h1]) + string(tobase64[h2])
// 		if o2 == 0 {
// 			enc += "=="
// 			break
// 		} else {
// 			enc += string(tobase64[h3])
// 		}
// 		if o3 == 0 {
// 			enc += "="
// 			break
// 		} else {
// 			enc += string(tobase64[h4])
// 		}
// 	}

// 	return enc

// }

func encode(data []byte) string {
	var result bytes.Buffer
	padding := 0

	for i := 0; i < len(data); i += 3 {
		var buffer uint32
		buffer = uint32(data[i]) << 16

		if i+1 < len(data) {
			buffer |= uint32(data[i+1]) << 8
		} else {
			padding++
		}

		if i+2 < len(data) {
			buffer |= uint32(data[i+2])
		} else {
			padding++
		}

		result.WriteByte(tobase64[(buffer>>18)&0x3F])
		result.WriteByte(tobase64[(buffer>>12)&0x3F])
		if padding < 2 {
			result.WriteByte(tobase64[(buffer>>6)&0x3F])
		} else {
			result.WriteByte('=')
		}
		if padding < 1 {
			result.WriteByte(tobase64[buffer&0x3F])
		} else {
			result.WriteByte('=')
		}
	}

	return result.String()
}

//shiza
// func decode(s string) ([]byte, error) {
// 	var result []byte
// 	for i := 0; i < len(s); i += 3 {
// 		c := s[i]

// 		if int(c) >= len(base64Map) {
// 			return nil, errors.New("this symbol is not a valid")
// 		}

// 		h := base64Map[c]

// 		if h == 255 {
// 			return nil, errors.New("this symbol is not a valid")
// 		}

// 		c = s[i+1]

// 		if int(c) >= len(base64Map) {
// 			return nil, errors.New("this symbol is not a valid")
// 		}

// 		l := base64Map[c]

// 		if l == 255 {
// 			return nil, errors.New("this symbol is not a valid")
// 		}

// 		c = s[i+2]

// 		if int(c) >= len(base64Map) {
// 			return nil, errors.New("this symbol is not a valid")
// 		}

// 		k := base64Map[c]

// 		if l == 255 {
// 			return nil, errors.New("this symbol is not a valid")
// 		}

// 		result[i/3] = (h << 12) | (l << 6) | k
// 	}
// 	return result, nil
// }

// func decode(s string) ([]byte, error) {
// 	result := make([]byte, len(s)*3/4)

// 	var resultIndex int
// 	for i := 0; i < len(s); i += 4 {
// 		if i+3 >= len(s) {
// 			return nil, errors.New("invalid base64 string length")
// 		}

// 		c1, c2, c3, c4 := s[i], s[i+1], s[i+2], s[i+3]

// 		v1, v2, v3, v4 := base64Map[c1], base64Map[c2], base64Map[c3], base64Map[c4]
// 		// if v1 == 255 || v2 == 255 || v3 == 255 || v4 == 255 {
// 		// 	return nil, errors.New("invalid character in base64 string")
// 		// }

// 		decodedByte1 := (v1 << 2) | (v2 >> 4)
// 		decodedByte2 := ((v2 & 0x0F) << 4) | (v3 >> 2)
// 		decodedByte3 := ((v3 & 0x03) << 6) | v4

// 		result[resultIndex] = decodedByte1
// 		resultIndex++
// 		if c3 != '=' {
// 			result[resultIndex] = decodedByte2
// 			resultIndex++
// 		}
// 		if c4 != '=' {
// 			result[resultIndex] = decodedByte3
// 			resultIndex++
// 		}
// 	}

// 	return result[:resultIndex], nil
// }

//my old method that wasnt work because of out of index issue
// func Decode(data string) ([]byte, error) {
// 	var output bytes.Buffer
// 	var chr1, chr2, chr3 byte
// 	var enc1, enc2, enc3, enc4 int
// 	i := 0

// 	data = cleanInput(data)

// 	for i < len(data) {
// 		enc1 = strings.IndexByte(tobase64, data[i])
// 		i++
// 		enc2 = strings.IndexByte(tobase64, data[i])
// 		i++
// 		enc3 = strings.IndexByte(tobase64, data[i])
// 		i++
// 		enc4 = strings.IndexByte(tobase64, data[i])
// 		i++

// 		if enc1 == -1 || enc2 == -1 || enc3 == -1 || enc4 == -1 {
// 			return []byte{}, errors.New("invalid Base64 data")
// 		}

// 		chr1 = byte((enc1 << 2) | (enc2 >> 4))
// 		chr2 = byte(((enc2 & 15) << 4) | (enc3 >> 2))
// 		chr3 = byte(((enc3 & 3) << 6) | enc4)

// 		output.WriteByte(chr1)

// 		if enc3 != 64 {
// 			output.WriteByte(chr2)
// 		}
// 		if enc4 != 64 {
// 			output.WriteByte(chr3)
// 		}
// 	}

// 	return output.Bytes(), nil
// }

func decode(data string) ([]byte, error) {
	var output bytes.Buffer
	var chr1, chr2, chr3 byte
	var enc1, enc2, enc3, enc4 byte
	i := 0

	data = delpadd.RemovePadding(data)

	for i < len(data) {
		if i >= len(data) {
			break
		}
		enc1 = base64Map[data[i]]
		i++

		if i >= len(data) {
			break
		}
		enc2 = base64Map[data[i]]
		i++

		enc3 = 64
		if i < len(data) {
			enc3 = base64Map[data[i]]
			i++
		}

		enc4 = 64
		if i < len(data) {
			enc4 = base64Map[data[i]]
			i++
		}

		if enc1 == 255 || enc2 == 255 || enc3 == 255 || enc4 == 255 {
			return nil, errors.New("invalid Base64 data")
		}

		chr1 = byte((enc1 << 2) | (enc2 >> 4))
		output.WriteByte(chr1)

		if enc3 != 64 {
			chr2 = byte(((enc2 & 15) << 4) | (enc3 >> 2))
			output.WriteByte(chr2)
		}

		if enc4 != 64 {
			chr3 = byte(((enc3 & 3) << 6) | enc4)
			output.WriteByte(chr3)
		}
	}

	return output.Bytes(), nil
}
