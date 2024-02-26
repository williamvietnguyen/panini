package core

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

// Position represents the state of a position
type Position struct {
	Pieces [NumPieceTypes]Bitboard
	Color  [NumColors]Bitboard
}

func (p *Position) GetColorOccupying(sq Square) Color {
	if p.Color[White].TestBit(sq) {
		return White
	}
	if p.Color[Black].TestBit(sq) {
		return Black
	}
	return NoColor
}

func (p *Position) GetPieceOnSquare(sq Square) Piece {
	pieceColor := p.GetColorOccupying(sq)
	if pieceColor == NoColor {
		return NoPiece
	}
	var pieceType PieceType
	for i, bb := range p.Pieces {
		if bb.TestBit(sq) {
			pieceType = int8(i)
		}
	}
	return Piece{Color: pieceColor, Type: pieceType}
}

func (p *Position) String() string {
	s := ""
	for rank := RankEight; rank >= RankOne; rank-- {
		s += fmt.Sprintf("%d", rank+1)
		for file := FileA; file <= FileH; file++ {
			sq := GetSquareNumber(rank, file)
			piece := p.GetPieceOnSquare(sq)
			if piece == NoPiece {
				s += "  ."
				continue
			}
			s += fmt.Sprintf("  %s", UnicodePieces[piece.Color][piece.Type])
		}
		s += "\n"
	}
	s += "   a  b  c  d  e  f  g  h\n"
	return s
}

func GetStartingPosition() Position {
	p := Position{
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
	return p
}
