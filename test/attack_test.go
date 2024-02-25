package test

import (
	"panini/internal"
	"testing"
)

func TestAttackMasks(t *testing.T) {
	a := internal.AttackMasks{}
	a.InitializeAttackTables()

	// white pawns
	if a.PawnAttacks[internal.White][63] != internal.Bitboard(0) {
		t.Fail()
	}
	if a.PawnAttacks[internal.White][56] != internal.Bitboard(0) {
		t.Fail()
	}
	if a.PawnAttacks[internal.White][0] != internal.Bitboard(0x0000000000000200) {
		t.Fail()
	}
	if a.PawnAttacks[internal.White][7] != internal.Bitboard(0x0000000000004000) {
		t.Fail()
	}
	if a.PawnAttacks[internal.White][28] != internal.Bitboard(0x0000002800000000) {
		t.Fail()
	}

	// black pawns
	if a.PawnAttacks[internal.Black][63] != internal.Bitboard(0x0040000000000000) {
		t.Fail()
	}
	if a.PawnAttacks[internal.Black][56] != internal.Bitboard(0x0002000000000000) {
		t.Fail()
	}
	if a.PawnAttacks[internal.Black][0] != internal.Bitboard(0) {
		t.Fail()
	}
	if a.PawnAttacks[internal.Black][7] != internal.Bitboard(0) {
		t.Fail()
	}
	if a.PawnAttacks[internal.Black][28] != internal.Bitboard(0x0000000000280000) {
		t.Fail()
	}

	// knight
	if a.KnightAttacks[63] != internal.Bitboard(0x0020400000000000) {
		t.Fail()
	}
	if a.KnightAttacks[56] != internal.Bitboard(0x0004020000000000) {
		t.Fail()
	}
	if a.KnightAttacks[0] != internal.Bitboard(0x0000000000020400) {
		t.Fail()
	}
	if a.KnightAttacks[7] != internal.Bitboard(0x0000000000402000) {
		t.Fail()
	}
	if a.KnightAttacks[22] != internal.Bitboard(0x000000A0100010A0) {
		t.Fail()
	}
	if a.KnightAttacks[33] != internal.Bitboard(0x0005080008050000) {
		t.Fail()
	}

	// king
	if a.KingAttacks[63] != internal.Bitboard(0x40C0000000000000) {
		t.Fail()
	}
	if a.KingAttacks[56] != internal.Bitboard(0x0203000000000000) {
		t.Fail()
	}
	if a.KingAttacks[0] != internal.Bitboard(0x0000000000000302) {
		t.Fail()
	}
	if a.KingAttacks[7] != internal.Bitboard(0x000000000000C040) {
		t.Fail()
	}
	if a.KingAttacks[22] != internal.Bitboard(0x00000000E0A0E000) {
		t.Fail()
	}
}
