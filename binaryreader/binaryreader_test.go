package binaryreader_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/felipe-bonato/goutils/binaryreader"
)

func TestPeek(t *testing.T) {
	buf := []byte{5, 6, 7}
	r := bytes.NewReader(buf)

	testPeekByte := func(expected byte, expectedOk bool) {
		b, ok := binaryreader.Peek[byte](r, binaryreader.LittleEndian)
		if ok != expectedOk {
			t.Errorf("expectedi ok to be %t, but got %t", expectedOk, ok)
		}
		if b != expected {
			t.Errorf("expected %d, but got %d", expected, b)
		}
	}

	testPeekByte(5, true)
	testPeekByte(5, true)

	// Advance reader
	r.Seek(1, io.SeekCurrent)
	testPeekByte(6, true)

	r.Seek(1, io.SeekCurrent)
	testPeekByte(7, true)

	r.Seek(1, io.SeekCurrent)
	testPeekByte(0, false)
}

func TestBinaryReaderByte(t *testing.T) {
	br := binaryreader.NewFromByteSlice([]byte{12, 13, 14}, binaryreader.LittleEndian)

	testReadByte := func(expectedByte byte, expectedOk bool) {
		b, ok := br.Byte()
		if ok != expectedOk {
			t.Errorf("read: expected ok to be %t, but got %t", expectedOk, ok)
		}
		if b != expectedByte {
			t.Errorf("read: expected %d, but got %d", expectedByte, b)
		}
	}

	testPeekByte := func(expectedByte byte, expectedOk bool) {
		b, ok := br.PeekByte()
		if ok != expectedOk {
			t.Errorf("peek: expected ok to be %t, but got %t", expectedOk, ok)
		}
		if b != expectedByte {
			t.Errorf("peek: expected %d, but got %d", expectedByte, b)
		}
	}

	testPeekByte(12, true)
	testPeekByte(12, true)
	testReadByte(12, true)

	testReadByte(13, true)

	testPeekByte(14, true)
	testReadByte(14, true)

	testReadByte(0, false)
	testPeekByte(0, false)
}

func TestBinaryReaderInt32LE(t *testing.T) {
	testBinaryReaderInt32(t, []byte{12, 0, 0, 0, 13, 0, 0, 0, 14, 0, 0, 0}, binaryreader.LittleEndian)
}

func TestBinaryReaderInt32BE(t *testing.T) {
	testBinaryReaderInt32(t, []byte{0, 0, 0, 12, 0, 0, 0, 13, 0, 0, 0, 14}, binaryreader.BigEndian)
}

func testBinaryReaderInt32(t *testing.T, data []byte, endianess binaryreader.Endianess) {
	br := binaryreader.New(bytes.NewReader(data), endianess)

	testReadInt := func(expectedInt int32, expectedOk bool) {
		i, ok := br.Int32()
		if ok != expectedOk {
			t.Errorf("read: expected ok to be %t, but got %t", expectedOk, ok)
		}
		if i != expectedInt {
			t.Errorf("read: expected %d, but got %d", expectedInt, i)
		}
	}

	testPeekInt := func(expectedInt int32, expectedOk bool) {
		i, ok := br.PeekInt32()
		if ok != expectedOk {
			t.Errorf("read: expected ok to be %t, but got %t", expectedOk, ok)
		}
		if i != expectedInt {
			t.Errorf("read: expected %d, but got %d", expectedInt, i)
		}
	}

	testPeekInt(12, true)
	testPeekInt(12, true)
	testReadInt(12, true)

	testReadInt(13, true)

	testPeekInt(14, true)
	testReadInt(14, true)

	testReadInt(0, false)
	testPeekInt(0, false)
}
