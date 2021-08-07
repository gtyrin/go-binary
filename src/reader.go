package binary

import (
	"bytes"
	"encoding/binary"
	"io"

	e "github.com/ytsiuryn/go-error"
)

// Reader хранит общие данные для универсального читателя потока с позиционированием.
type Reader struct {
	r   io.ReadSeeker
	buf []byte
}

// NewReader создает новый объект типа Reader.
func NewReader(r io.ReadSeeker) *Reader {
	return &Reader{r: r, buf: make([]byte, 256)}
}

// Position return the current file position
func (b *Reader) Position() int64 {
	return b.SeekBytes(0, io.SeekCurrent)
}

// SeekBytes moves file curcsor
func (b *Reader) SeekBytes(n int64, direction int) int64 {
	pos, err := b.r.Seek(n, direction)
	e.Panic(err, "SeekBytes")
	return pos
}

// CheckBytes reads bytes in the current position without file cursor moving
func (b *Reader) CheckBytes(nBytes int64) []byte {
	n, err := b.r.Read(b.buf[:nBytes])
	e.Panic(err, "CheckBytes")
	_, err = b.r.Seek(-int64(n), io.SeekCurrent)
	e.Panic(err, "CheckBytes")
	return b.buf[:n]
}

// ReadBytes reads nBytes bytes from file
func (b *Reader) ReadBytes(n int64) []byte {
	if int(n) > len(b.buf) {
		b.buf = make([]byte, n)
	}
	n2, err := b.r.Read(b.buf[:n])
	e.Panic(err, "ReadBytes")
	return b.buf[:n2]
	// bytes := make([]byte, n)
	// _, err := f.Read(bytes)
	// system.CheckError(err, "ReadBytes")
	// return bytes
}

// ReadInto reads bytes into some structure.
func (b *Reader) ReadInto(n int64, order binary.ByteOrder, out interface{}) {
	data := b.ReadBytes(n)
	buf := bytes.NewBuffer(data)
	err := binary.Read(buf, order, out)
	e.Panic(err, "ReadInto")
}

// SkipBytes `n` bytes from the cureent position.
func (b *Reader) SkipBytes(n int64) int64 {
	pos, err := b.r.Seek(n, io.SeekCurrent)
	e.Panic(err, "SkipBytes")
	return pos
}

// SkipNullBytes skips null bytes
func (b *Reader) SkipNullBytes() int64 {
	var skippedBytes int64
	zero := []byte{0}
	for bytes.Equal(b.ReadBytes(1), zero) {
		skippedBytes++
	}
	b.SeekBytes(-1, io.SeekCurrent)
	return skippedBytes - 1
}

// ReadString reads bytes upto 0x00 byte
func (b *Reader) ReadString() string {
	var err error
	i := 0
	for {
		if i == len(b.buf) {
			b.buf = make([]byte, 2*len(b.buf))
		}
		_, err = b.r.Read(b.buf[i : i+1])
		e.Panic(err, "ReadString")
		if b.buf[i] == 0 {
			break
		}
		i++
	}
	return string(b.buf[:i])
}

// ReadUint8 is read helper for uint8 value with a panic if error exists.
func (b *Reader) ReadUint8() uint8 {
	e.Panic(binary.Read(b.r, binary.BigEndian, b.buf[:1]), "ReadUint8")
	return b.buf[0]
}

// ReadBEUint16 is read helper for uint16 value with a panic if error exists.
func (b *Reader) ReadBEUint16() uint16 {
	var res uint16
	e.Panic(binary.Read(b.r, binary.BigEndian, &res), "ReadBEUint16")
	return res
}

// ReadBEInt32 is read helper for int32 value with a panic if error exists.
func (b *Reader) ReadBEInt32() int32 {
	var res int32
	e.Panic(binary.Read(b.r, binary.BigEndian, &res), "ReadBEInt32")
	return res
}

// ReadBEUint32 is read helper for uint32 value with a panic if error exists.
func (b *Reader) ReadBEUint32() uint32 {
	var res uint32
	e.Panic(binary.Read(b.r, binary.BigEndian, &res), "ReadBEUint32")
	return res
}

// ReadBEInt64 is read helper for int64 value with a panic if error exists.
func (b *Reader) ReadBEInt64() int64 {
	var res int64
	e.Panic(binary.Read(b.r, binary.BigEndian, &res), "ReadBEInt64")
	return res
}

// ReadBEUint64 is read helper for uint64 value with a panic if error exists.
func (b *Reader) ReadBEUint64() uint64 {
	var res uint64
	e.Panic(binary.Read(b.r, binary.BigEndian, &res), "ReadBEUint64")
	return res
}

// ReadLEUint16 is read helper for uint16 value with a panic if error exists.
func (b *Reader) ReadLEUint16() uint16 {
	var res uint16
	e.Panic(binary.Read(b.r, binary.LittleEndian, &res), "ReadLEUint16")
	return res
}

// ReadLEInt32 is read helper for uint32 value with a panic if error exists.
func (b *Reader) ReadLEInt32() int32 {
	var res int32
	e.Panic(binary.Read(b.r, binary.LittleEndian, &res), "ReadLEInt32")
	return res
}

// ReadLEUint32 is read helper for uint32 value with a panic if error exists.
func (b *Reader) ReadLEUint32() uint32 {
	var res uint32
	e.Panic(binary.Read(b.r, binary.LittleEndian, &res), "ReadLEUint32")
	return res
}

// ReadLEInt64 is read helper for int64 value with a panic if error exists.
func (b *Reader) ReadLEInt64() int64 {
	var res int64
	e.Panic(binary.Read(b.r, binary.LittleEndian, &res), "ReadLEInt64")
	return res
}

// ReadLEUint64 is read helper for uint64 value with a panic if error exists.
func (b *Reader) ReadLEUint64() uint64 {
	var res uint64
	e.Panic(binary.Read(b.r, binary.LittleEndian, &res), "ReadLEUint64")
	return res
}
