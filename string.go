package binary

import (
	"fmt"
	"unicode/utf8"

	"golang.org/x/text/encoding/unicode"
)

// FromASCIIZ cuts all padding zero bytes.
func FromASCIIZ(b []byte) []byte {
	i := len(b) - 1
	for ; i > 0 && b[i] == 0; i-- {
	}
	return b[:i+1]
}

func fromUTF16(b []byte, order unicode.Endianness) ([]byte, error) {
	return unicode.UTF16(order, unicode.ExpectBOM).NewDecoder().Bytes(b)
}

// FromUTF16LE converts string from UTF16 little endian to UTF8.
func FromUTF16LE(b []byte) ([]byte, error) {
	return fromUTF16(b, unicode.LittleEndian)
}

// FromUTF16BE converts string from UTF16 big endian to UTF8.
func FromUTF16BE(b []byte) ([]byte, error) {
	ret, err := fromUTF16(b, unicode.BigEndian)
	fmt.Printf("big endian: b=%#v, result=%s\n", b, ret)
	return ret, err
}

// PanicIfNonUtf8 panics if the input string is non UTF-8 one.
func PanicIfNonUtf8(strings ...string) {
	for _, s := range strings {
		if !utf8.ValidString(s) {
			panic("non-UTF8 string: " + s)
		}
	}
}
