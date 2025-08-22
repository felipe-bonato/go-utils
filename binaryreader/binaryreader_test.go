package binaryreader_test

import (
	"bytes"
	"testing"

	"github.com/felipe-bonato/goutils/binaryreader"
)

func TestBinaryReaderByte(t *testing.T) {
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

func TestBinaryReaderInt(t *testing.T) {
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
