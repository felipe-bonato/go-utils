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

func Read[T any](reader io.Reader, endianess Endianess) (T, bool) {
	var x T
	return x, binary.Read(reader, endianess, &x) == nil
}

// Reads the next `count` bytes into an array.
func (br *BinaryReader) Bytes(count int) ([]byte, bool) {
	bytes := []byte{}

	for range count {
		b, ok := br.Byte()
		if !ok {
			return bytes, false
		}

		bytes = append(bytes, b)
	}

	return bytes, true
}
