package binaryreader

import (
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
	r         io.Reader
	Endianess binary.ByteOrder
}

func New(r io.Reader, endianess Endianess) BinaryReader {
	return BinaryReader{r, endianess}
}

func (br *BinaryReader) Byte() (byte, bool) { return Read[byte](br.r, br.Endianess) }

func (br *BinaryReader) Int8() (int8, bool)   { return Read[int8](br.r, br.Endianess) }
func (br *BinaryReader) Int16() (int16, bool) { return Read[int16](br.r, br.Endianess) }
func (br *BinaryReader) Int32() (int32, bool) { return Read[int32](br.r, br.Endianess) }
func (br *BinaryReader) Int64() (int64, bool) { return Read[int64](br.r, br.Endianess) }

func (br *BinaryReader) Uint8() (uint8, bool)   { return Read[uint8](br.r, br.Endianess) }
func (br *BinaryReader) Uint16() (uint16, bool) { return Read[uint16](br.r, br.Endianess) }
func (br *BinaryReader) Uint32() (uint32, bool) { return Read[uint32](br.r, br.Endianess) }
func (br *BinaryReader) Uint64() (uint64, bool) { return Read[uint64](br.r, br.Endianess) }

func (br *BinaryReader) Float32() (float32, bool) { return Read[float32](br.r, br.Endianess) }
func (br *BinaryReader) Float64() (float64, bool) { return Read[float64](br.r, br.Endianess) }

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
