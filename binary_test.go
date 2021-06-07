package binary

import (
	"bytes"
	"testing"
)

var testData = []byte{}

func TestFromASCIIZ(t *testing.T) {
	b := []byte{0x54, 0x65, 0x73, 0x74, 0x0, 0x0, 0x0, 0x0}
	res := FromASCIIZ(b)
	if string(res) != "Test" {
		t.Fail()
	}
}

func TestFromUTF16LE(t *testing.T) {
	b := []byte{
		0xff, 0xfe, 0x14, 0x4, 0x32, 0x4, 0x30, 0x4, 0x20, 0x0, 0x3a, 0x4, 0x3e, 0x4, 0x3b, 0x4,
		0x4c, 0x4, 0x46, 0x4, 0x30, 0x4}
	if res, err := FromUTF16LE(b); err != nil || string(res) != "Два кольца" {
		t.Fail()
	}
}

func TestFromUTF16BE(t *testing.T) {

}

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
