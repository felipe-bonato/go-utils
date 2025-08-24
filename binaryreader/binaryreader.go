package binaryreader

import (
	"bytes"
	"encoding/binary"
	"io"
)

type Endianess = binary.ByteOrder

var (
	LittleEndian Endianess = binary.LittleEndian
	BigEndian    Endianess = binary.BigEndian
	NativeEndian Endianess = binary.NativeEndian
)

// Reads `T` from `r` with `endianess`. If `r` ended, returns `_, false`
//
// `T` must have compile-time known size. If not, will always returns `_, false`
func Read[T any](r io.Reader, endianess Endianess) (T, bool) {
	var x T
	return x, binary.Read(r, endianess, &x) == nil
}

// Reads `T` from `r` with `endianess`, but doesn't advance it. If `r` ended, returns `_, false`
//
// `T` must have compile-time known size. If not, will always returns `_, false`
func Peek[T any](rs io.ReadSeeker, endianess Endianess) (T, bool) {
	startPoint, err := rs.Seek(0, io.SeekCurrent)
	if err != nil {
		panic("saving seeking point failed")
	}

	x, ok := Read[T](rs, endianess)

	// Reset the reader
	_, err = rs.Seek(startPoint, io.SeekStart)
	if err != nil {
		panic("reseting seeking point failed")
	}

	return x, ok
}

type BinaryReader struct {
	rs        io.ReadSeeker
	Endianess binary.ByteOrder
}

func New(rs io.ReadSeeker, endianess Endianess) BinaryReader {
	return BinaryReader{rs, endianess}
}

// IMPORTANT: when creating with this function, the peek functionality doesn't work.
// Calling it will panic. Use the New instead.
func NewFromReader(r io.Reader, endianess Endianess) BinaryReader {
	return BinaryReader{&readSeekfier{r}, endianess}
}

func NewFromByteSlice(s []byte, endianess Endianess) BinaryReader {
	return BinaryReader{bytes.NewReader(s), endianess}
}

func (br *BinaryReader) Byte() (byte, bool) { return Read[byte](br.rs, br.Endianess) }

func (br *BinaryReader) Int8() (int8, bool)   { return Read[int8](br.rs, br.Endianess) }
func (br *BinaryReader) Int16() (int16, bool) { return Read[int16](br.rs, br.Endianess) }
func (br *BinaryReader) Int32() (int32, bool) { return Read[int32](br.rs, br.Endianess) }
func (br *BinaryReader) Int64() (int64, bool) { return Read[int64](br.rs, br.Endianess) }

func (br *BinaryReader) Uint8() (uint8, bool)   { return Read[uint8](br.rs, br.Endianess) }
func (br *BinaryReader) Uint16() (uint16, bool) { return Read[uint16](br.rs, br.Endianess) }
func (br *BinaryReader) Uint32() (uint32, bool) { return Read[uint32](br.rs, br.Endianess) }
func (br *BinaryReader) Uint64() (uint64, bool) { return Read[uint64](br.rs, br.Endianess) }

func (br *BinaryReader) Float32() (float32, bool) { return Read[float32](br.rs, br.Endianess) }
func (br *BinaryReader) Float64() (float64, bool) { return Read[float64](br.rs, br.Endianess) }

// Reads the next `count` bytes into an array.
func (br *BinaryReader) Bytes(count int) ([]byte, bool) {
	// TODO: Directly reading from r is probably faster.
	bytes := make([]byte, 0, count)

	for range count {
		b, ok := br.Byte()
		if !ok {
			return bytes, false
		}

		bytes = append(bytes, b)
	}

	return bytes, true
}

// TODO: ReadStringNullTerminated
// TODO: ReadStringSized

func (br *BinaryReader) PeekByte() (byte, bool) { return Peek[byte](br.rs, br.Endianess) }

func (br *BinaryReader) PeekInt8() (int8, bool)   { return Peek[int8](br.rs, br.Endianess) }
func (br *BinaryReader) PeekInt16() (int16, bool) { return Peek[int16](br.rs, br.Endianess) }
func (br *BinaryReader) PeekInt32() (int32, bool) { return Peek[int32](br.rs, br.Endianess) }
func (br *BinaryReader) PeekInt64() (int64, bool) { return Peek[int64](br.rs, br.Endianess) }

func (br *BinaryReader) PeekUint8() (uint8, bool)   { return Peek[uint8](br.rs, br.Endianess) }
func (br *BinaryReader) PeekUint16() (uint16, bool) { return Peek[uint16](br.rs, br.Endianess) }
func (br *BinaryReader) PeekUint32() (uint32, bool) { return Peek[uint32](br.rs, br.Endianess) }
func (br *BinaryReader) PeekUint64() (uint64, bool) { return Peek[uint64](br.rs, br.Endianess) }

func (br *BinaryReader) PeekFloat32() (float32, bool) { return Peek[float32](br.rs, br.Endianess) }
func (br *BinaryReader) PeekFloat64() (float64, bool) { return Peek[float64](br.rs, br.Endianess) }

// TODO: PeekBytes
// TODO: PeekStringNullTerminated
// TODO: PeekStringSized

// TODO: Advance

type readSeekfier struct {
	io.Reader
}

func (rs *readSeekfier) Seek(_ int64, _ int) (int64, error) {
	panic("not implemented for type")
}
