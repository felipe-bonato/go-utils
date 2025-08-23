package binaryreader_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/felipe-bonato/goutils/binaryreader"
)

func TestReadByte(t *testing.T) {
	r := bytes.NewReader([]byte{12, 13, 14})

	br := binaryreader.New(r, binaryreader.LittleEndian)

	testReadByte := func(expected byte) {
		b, ok := br.Byte()
		if !ok || b != expected {
			t.Errorf("expected %v, but got %v", expected, b)
		}
	}

	testReadByte(12)
	testReadByte(13)
	testReadByte(14)
	if _, ok := br.Byte(); ok { // Ended, should return false
		t.Fatal("Should have ended")
	}

}

func TestReadInt32(t *testing.T) {
	buf := []byte{12, 0, 0, 0, 13, 0, 0, 0, 14, 0, 0, 0}
	br := binaryreader.New(bytes.NewReader(buf), binaryreader.LittleEndian)

	testReadInt := func(expected int32) {
		i, ok := br.Int32()
		if !ok || i != expected {
			t.Errorf("expected %v, but got %v", expected, i)
		}
	}

	testReadInt(12)
	testReadInt(13)
	testReadInt(14)
	if _, ok := br.Int32(); ok { // Ended, should return false
		t.Fatal("Should have ended")
	}

}

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
