package binaryreader

import (
	"encoding/binary"
	"io"
)

type BinaryReader struct {
	r         io.Reader
	Endianess binary.ByteOrder
}

type Endianess = binary.ByteOrder

var LittleEndian Endianess = binary.LittleEndian
var BigEndian Endianess = binary.BigEndian
var NativeEndian Endianess = binary.NativeEndian

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

// Reads `T` from `r` with `endianess`. If `r` ended, return `_, false`
//
// `T` must have compile-time known size. If not, will always return `_, false`
func Read[T any](r io.Reader, endianess Endianess) (T, bool) {
	var x T
	return x, binary.Read(r, endianess, &x) == nil
}

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

func Peek[T any](rs io.ReadSeeker, endianess Endianess) (T, bool) {
	currReadIndex, err := rs.Seek(0, io.SeekCurrent)
	if err != nil {
		panic("saving seeking index failed")
	}

	x, ok := Read[T](rs, endianess)

	// Reset the reader
	_, err = rs.Seek(currReadIndex, io.SeekStart)
	if err != nil {
		panic("reseting seeking index failed")
	}

	return x, ok
}
