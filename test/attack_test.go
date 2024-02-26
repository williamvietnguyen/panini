package test

import (
	"github.com/williamvietnguyen/panini/internal/core"
	"testing"
)

func TestAttackMasks(t *testing.T) {
	core.InitializeAttackTables()

	// white pawns
	if core.PawnAttacks[core.White][63] != core.Bitboard(0) {
		t.Fail()
	}
	if core.PawnAttacks[core.White][56] != core.Bitboard(0) {
		t.Fail()
	}
	if core.PawnAttacks[core.White][0] != core.Bitboard(0x0000000000000200) {
		t.Fail()
	}
	if core.PawnAttacks[core.White][7] != core.Bitboard(0x0000000000004000) {
		t.Fail()
	}
	if core.PawnAttacks[core.White][28] != core.Bitboard(0x0000002800000000) {
		t.Fail()
	}

	// black pawns
	if core.PawnAttacks[core.Black][63] != core.Bitboard(0x0040000000000000) {
		t.Fail()
	}
	if core.PawnAttacks[core.Black][56] != core.Bitboard(0x0002000000000000) {
		t.Fail()
	}
	if core.PawnAttacks[core.Black][0] != core.Bitboard(0) {
		t.Fail()
	}
	if core.PawnAttacks[core.Black][7] != core.Bitboard(0) {
		t.Fail()
	}
	if core.PawnAttacks[core.Black][28] != core.Bitboard(0x0000000000280000) {
		t.Fail()
	}

	// knight
	if core.KnightAttacks[63] != core.Bitboard(0x0020400000000000) {
		t.Fail()
	}
	if core.KnightAttacks[56] != core.Bitboard(0x0004020000000000) {
		t.Fail()
	}
	if core.KnightAttacks[0] != core.Bitboard(0x0000000000020400) {
		t.Fail()
	}
	if core.KnightAttacks[7] != core.Bitboard(0x0000000000402000) {
		t.Fail()
	}
	if core.KnightAttacks[22] != core.Bitboard(0x000000A0100010A0) {
		t.Fail()
	}
	if core.KnightAttacks[33] != core.Bitboard(0x0005080008050000) {
		t.Fail()
	}

	// king
	if core.KingAttacks[63] != core.Bitboard(0x40C0000000000000) {
		t.Fail()
	}
	if core.KingAttacks[56] != core.Bitboard(0x0203000000000000) {
		t.Fail()
	}
	if core.KingAttacks[0] != core.Bitboard(0x0000000000000302) {
		t.Fail()
	}
	if core.KingAttacks[7] != core.Bitboard(0x000000000000C040) {
		t.Fail()
	}
	if core.KingAttacks[22] != core.Bitboard(0x00000000E0A0E000) {
		t.Fail()
	}
}
