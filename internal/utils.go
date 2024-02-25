package internal

type Square = int8
type Rank = int8
type File = int8

const (
	RankOne Rank = iota
	RankTwo Rank = iota
	RankThree Rank = iota
	RankFour Rank = iota
	RankFive Rank = iota
	RankSix Rank = iota
	RankSeven Rank = iota
	RankEight Rank = iota
)

const (
	FileA File = iota
	FileB File = iota
	FileC File = iota
	FileD File = iota
	FileE File = iota
	FileF File = iota
	FileG File = iota
	FileH File = iota
)
