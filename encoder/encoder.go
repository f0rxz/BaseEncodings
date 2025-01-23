package encoder

import (
	"mystdencodings/base16util"
	"mystdencodings/base32util"
	"mystdencodings/base64util"
)

type Encoder struct {
	Base64 *base64util.Base64Encoder
	Base32 *base32util.Base32Encoder
	Base16 *base16util.Base16Encoder
}

func NewEncoder() *Encoder {
	return &Encoder{
		Base64: base64util.NewEncoder(),
		Base32: base32util.NewEncoder(),
		Base16: base16util.NewEncoder(),
	}
}
