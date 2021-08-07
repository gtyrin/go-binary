package binary

import (
	"bytes"
	"testing"
)

var testData = []byte{}

// TODO: читать двоичный образец и пройтись по всем функциям в одной этой тестовой функции.
func TestBinaryReader(t *testing.T) {
	r := NewReader(bytes.NewReader(testData))
	r.Position()
	// r.SeekBytes
	// r.CheckBytes
	// r.ReadBytes
	// r.ReadInto
	// r.SkipBytes
	// r.SkipNullBytes
	// r.ReadString
	// r.ReadUint8
	// r.ReadBEUint16
	// r.ReadBEInt32
	// r.ReadBEUint32
	// r.ReadBEInt64
	// r.ReadBEUint64
	// r.ReadLEUint16
	// r.ReadLEInt32
	// r.ReadLEUint32
	// r.ReadLEInt64
	// r.ReadLEUint64
}
