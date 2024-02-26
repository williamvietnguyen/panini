package internal

import "fmt"

// Bitboard is an unsigned 64-bit integer where the
// least significant bit corresponds with a1 and
// most significant bit corresponds with h8
type Bitboard uint64

func (bb *Bitboard) TestBit(sq Square) bool {
	return (*bb>>sq)&Bitboard(1) == 1
}

func (bb *Bitboard) SetBit(sq Square) {
	*bb |= Bitboard(1) << sq
}

func (bb *Bitboard) PopBit(sq Square) {
	if bb.TestBit(sq) {
		*bb ^= Bitboard(1) << sq
	}
}

func GetSquareNumber(r Rank, f File) Square {
	return r<<3 + f
}

func (bb *Bitboard) String() string {
	s := ""
	for rank := RankEight; rank >= RankOne; rank-- {
		s += fmt.Sprintf("%d", rank+1)
		for file := FileA; file <= FileH; file++ {
			if bb.TestBit(GetSquareNumber(rank, file)) {
				s += "  x"
			} else {
				s += "  ."
			}
		}
		s += "\n"
	}
	s += "   a  b  c  d  e  f  g  h\n"
	s += fmt.Sprintf("Bitboard: 0x%016X\n", *bb)
	return s
}
