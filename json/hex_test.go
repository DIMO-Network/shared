package json

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestMarshalJSON(t *testing.T) {
	cases := []struct {
		Name     string
		Slice    HexBytes
		Expected string
	}{
		{
			Name:     "Nil slice",
			Slice:    nil,
			Expected: "null",
		},
		{
			Name:     "Zero-length slice",
			Slice:    []byte{},
			Expected: `"0x"`,
		},
		{
			Name:     "Typical slice",
			Slice:    []byte{26, 43, 60},
			Expected: `"0x1a2b3c"`,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			b, err := json.Marshal(HexBytes(c.Slice))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			actual := string(b)

			if c.Expected != actual {
				t.Errorf("expected %q, got %q", c.Expected, actual)
			}
		})
	}
}

func TestUnmarshalJSONNull(t *testing.T) {
	var b HexBytes
	err := json.Unmarshal([]byte(`null`), &b)
	if err != nil {
		t.Fatalf("unexpected error; %v", err)
	}

	if b != nil {
		t.Errorf("expected slice to be nil, but had values %s", b)
	}
}

func TestUnmarshalJSONSuccess(t *testing.T) {
	cases := []struct {
		Name     string
		JSON     string
		Expected []byte
	}{
		{
			Name:     "No characters",
			JSON:     `""`,
			Expected: []byte{},
		},
		{
			Name:     "No hex digits",
			JSON:     `"0x"`,
			Expected: []byte{},
		},
		{
			Name:     "Normal with 0x prefix",
			JSON:     `"0x1a2b3c"`,
			Expected: []byte{26, 43, 60},
		},
		{
			Name:     "Normal with 0X prefix",
			JSON:     `"0X1a2b3c"`,
			Expected: []byte{26, 43, 60},
		},
		{
			Name:     "Normal with no prefix",
			JSON:     `"1a2b3c"`,
			Expected: []byte{26, 43, 60},
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			var b HexBytes
			err := json.Unmarshal([]byte(c.JSON), &b)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !bytes.Equal(c.Expected, b) {
				t.Errorf("expected %s, got %s", c.Expected, b)
			}
		})
	}
}

func TestUnmarshalJSONFailure(t *testing.T) {
	cases := []struct {
		Name  string
		JSON  string
		Error string
	}{
		{
			Name: "Odd number of hex digits",
			JSON: `"0x1a2b3"`,
		},
		{
			Name: "JSON number",
			JSON: `45`,
		},
		{
			Name: "Invalid prefix",
			JSON: `"0Y1a2b3c"`,
		},
		{
			Name: "Non-hexadecimal digits",
			JSON: `"0x4r"`,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			var b HexBytes
			err := json.Unmarshal([]byte(c.JSON), &b)
			if err == nil {
				t.Errorf("expected an error, but got value %s and no error", b)
			}
		})
	}
}
