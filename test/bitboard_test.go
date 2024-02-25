package test

import (
	"panini/internal"
	"testing"
)

func TestSetBit(t *testing.T) {
	bb := internal.Bitboard(1)
	bb.SetBit(1)
	if bb != internal.Bitboard(3) {
		t.Fail()
	}
	bb.SetBit(63)
	if bb != internal.Bitboard(0x8000000000000003) {
		t.Fail()
	}
}

func TestPopBit(t *testing.T) {
	bb := internal.Bitboard(0x8000000000000003)
	bb.PopBit(1)
	if bb != internal.Bitboard(0x8000000000000001) {
		t.Error()
	}
	bb.PopBit(63)
	if bb != internal.Bitboard(1) {
		t.Fail()
	}
}

func TestTestBit(t *testing.T) {
	bb := internal.Bitboard(0x8000000000000003)
	if !bb.TestBit(1) {
		t.Fail()
	}
	if bb.TestBit(40) {
		t.Fail()
	}
}
