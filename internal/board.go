package internal

import "fmt"

type PieceType = int8
type Color = int8

type Piece struct {
	Color Color
	Type  PieceType
}

const (
	White   Color = iota
	Black   Color = iota
	NoColor Color = iota
)

const NumColors = 2
const NumPieceTypes = 6

const (
	Pawn       PieceType = iota
	Knight     PieceType = iota
	Bishop     PieceType = iota
	Rook       PieceType = iota
	Queen      PieceType = iota
	King       PieceType = iota
	EmptyPiece PieceType = iota
)

var NoPiece = Piece{Color: NoColor, Type: EmptyPiece}

type BoardState struct {
	Pieces [NumPieceTypes]Bitboard
	Color  [NumColors]Bitboard
}

func (bs *BoardState) GetColorOccupying(sq Square) Color {
	if bs.Color[White].TestBit(sq) {
		return White
	}
	if bs.Color[Black].TestBit(sq) {
		return Black
	}
	return NoColor
}

func (bs *BoardState) GetPieceOnSquare(sq Square) Piece {
	pieceColor := bs.GetColorOccupying(sq)
	if pieceColor == NoColor {
		return NoPiece
	}
	var pieceType PieceType
	for i, bb := range bs.Pieces {
		if bb.TestBit(sq) {
			pieceType = int8(i)
		}
	}
	return Piece{Color: pieceColor, Type: pieceType}
}

func (bs *BoardState) Print() {
	for rank := RankEight; rank >= RankOne; rank-- {
		fmt.Printf("%d", rank+1)
		for file := FileA; file <= FileH; file++ {
			sq := GetSquareNumber(rank, file)
			piece := bs.GetPieceOnSquare(sq)
			if piece == NoPiece {
				fmt.Printf("  .")
				continue
			}
			fmt.Printf("  %s", UnicodePieces[piece.Color][piece.Type])
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
