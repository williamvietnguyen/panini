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
	return 8*r + f
}

func (bb *Bitboard) Print() {
	for rank := RankEight; rank >= RankOne; rank-- {
		fmt.Printf("%d", rank+1)
		for file := FileA; file <= FileH; file++ {
			if bb.TestBit(GetSquareNumber(rank, file)) {
				fmt.Print("  x")
			} else {
				fmt.Print("  .")
			}
		}
		fmt.Println()
	}
	fmt.Println("   a  b  c  d  e  f  g  h")
	fmt.Printf("\nBitboard: 0x%016X", *bb)
}
