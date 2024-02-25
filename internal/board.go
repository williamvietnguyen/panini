package internal

import "fmt"

type Piece = int8
type Color = int8

const (
	White   Color = iota
	Black   Color = iota
	NoColor Color = iota
)

const NumColors = 2
const NumPieceTypes = 6

const (
	Pawn    Piece = iota
	Knight  Piece = iota
	Bishop  Piece = iota
	Rook    Piece = iota
	Queen   Piece = iota
	King    Piece = iota
	NoPiece Piece = iota
)

var UnicodePieces = [NumColors][NumPieceTypes]string{{"♙", "♘", "♗", "♖", "♕", "♔"}, {"♟", "♞", "♝", "♜", "♛", "♚"}}

type BoardState struct {
	Pieces [NumPieceTypes]Bitboard
	Color  [NumColors]Bitboard
}

func (bs *BoardState) GetPieceOnSquare(sq Square) Piece {
	if bs.Pieces[Pawn].TestBit(sq) {
		return Pawn
	}
	if bs.Pieces[Knight].TestBit(sq) {
		return Knight
	}
	if bs.Pieces[Bishop].TestBit(sq) {
		return Bishop
	}
	if bs.Pieces[Rook].TestBit(sq) {
		return Rook
	}
	if bs.Pieces[Queen].TestBit(sq) {
		return Queen
	}
	if bs.Pieces[King].TestBit(sq) {
		return King
	}
	return NoPiece
}

func (bs *BoardState) GetColorOnSquare(sq Square) Color {
	if bs.Color[White].TestBit(sq) {
		return White
	}
	if bs.Color[Black].TestBit(sq) {
		return Black
	}
	return NoColor
}

func (bs *BoardState) Print() {
	for rank := RankEight; rank >= RankOne; rank-- {
		fmt.Printf("%d", rank+1)
		for file := FileA; file <= FileH; file++ {
			sq := GetSquareNumber(rank, file)
			color := bs.GetColorOnSquare(sq)
			if color == NoColor {
				fmt.Printf("  .")
				continue
			}
			piece := bs.GetPieceOnSquare(sq)
			if piece == NoPiece {
				fmt.Printf("  .")
				continue
			}
			fmt.Printf("  %s", UnicodePieces[color][piece])
		}
		fmt.Println()
	}
	fmt.Println("   a  b  c  d  e  f  g  h")
}

func InitializeBoardState() BoardState {
	bs := BoardState{
		Pieces: [NumPieceTypes]Bitboard{
			0x00ff00000000ff00,
			0x4200000000000042,
			0x2400000000000024,
			0x8100000000000081,
			0x0800000000000008,
			0x1000000000000010,
		},
		Color: [NumColors]Bitboard{
			0x000000000000ffff,
			0xffff000000000000,
		},
	}
	return bs
}
