package main

import (
	"flag"
	"fmt"
	"os"

	"mystdencodings/encoder"
)

func main() {
	encodingFlag := flag.String("encoding", "base64", "Encoding for usage (base64, base32 or base16)")
	flag.Parse()
	enc := encoder.NewEncoder()

	input := "Hello, World!"
	inputBytes := []byte(input)

	var encoded string
	switch *encodingFlag {
	case "base64":
		encoded = enc.Base64.Encode(inputBytes)
	case "base32":
		encoded = enc.Base32.Encode(inputBytes)
	case "base16":
		encoded = enc.Base16.Encode(inputBytes)
	default:
		fmt.Fprintf(os.Stderr, "Unsupported encoding: %s\n", *encodingFlag)
		os.Exit(1)
	}

	fmt.Printf("Source string: %s\n", input)
	fmt.Printf("Encoded string (%s): %s\n", *encodingFlag, encoded)
}
