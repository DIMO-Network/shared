package json

import (
	"encoding/hex"
	"errors"
	"fmt"
)

type HexBytes []byte

func (b HexBytes) MarshalJSON() ([]byte, error) {
	if b == nil {
		return []byte("null"), nil
	}

	// Two bytes for the quotes, two bytes for the 0x prefix.
	outLen := 4 + hex.EncodedLen(len(b))

	out := make([]byte, outLen)
	copy(out, `"0x`)
	out[outLen-1] = '"'
	hex.Encode(out[3:], b)
	return out, nil
}

func (b *HexBytes) UnmarshalJSON(text []byte) error {
	if string(text) == "null" {
		return nil
	}

	if !isString(text) {
		return errors.New("marshaling a JSON value that is not null or a string")
	}

	text = text[1 : len(text)-1]

	if has0xPrefix(text) {
		text = text[2:]
	}

	if len(text)%2 != 0 {
		return errors.New("hex string had an odd length")
	}

	out := make([]byte, len(text)/2)

	_, err := hex.Decode(out, text)
	if err != nil {
		return fmt.Errorf("error decoding hex: %w", err)
	}

	*b = out
	return nil
}

func has0xPrefix(text []byte) bool {
	return len(text) >= 2 && text[0] == '0' && (text[1] == 'x' || text[1] == 'X')
}

func isString(text []byte) bool {
	return len(text) >= 2 && text[0] == '"' && text[len(text)-1] == '"'
}
