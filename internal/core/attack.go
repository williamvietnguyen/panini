package core

const NotHFile Bitboard = 0x7F7F7F7F7F7F7F7F
const NotAFile Bitboard = 0xFEFEFEFEFEFEFEFE
const NotABFile Bitboard = 0xFCFCFCFCFCFCFCFC
const NotGHFile Bitboard = 0x3F3F3F3F3F3F3F3F

var PawnAttacks = [NumColors][NumSquares]Bitboard{}
var KnightAttacks = [NumSquares]Bitboard{}
var KingAttacks = [NumSquares]Bitboard{}

func InitializeAttackTables() {
	for i := int8(0); i < NumSquares; i++ {
		var bb Bitboard
		bb.SetBit(i)

		// pawns
		PawnAttacks[White][i] = (bb & NotHFile << 9) | (bb & NotAFile << 7)
		PawnAttacks[Black][i] = (bb & NotAFile >> 9) | (bb & NotHFile >> 7)

		// knights
		KnightAttacks[i] = (bb & NotAFile << 15) | (bb & NotAFile >> 17) | (bb & NotABFile << 6) |
			(bb & NotGHFile << 10) | (bb & NotHFile << 17) | (bb & NotHFile >> 15) |
			(bb & NotABFile >> 10) | (bb & NotGHFile >> 6)

		// kings
		KingAttacks[i] |= PawnAttacks[White][i] | PawnAttacks[Black][i]
		KingAttacks[i] |= (bb & NotAFile >> 1) | (bb & NotHFile << 1) | (bb >> 8) | (bb << 8)
	}
}
