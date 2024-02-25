package internal

type Square = int8
type Rank = int8
type File = int8

const (
	RANK_1 Rank = iota
	RANK_2 Rank = iota
	RANK_3 Rank = iota
	RANK_4 Rank = iota
	RANK_5 Rank = iota
	RANK_6 Rank = iota
	RANK_7 Rank = iota
	RANK_8 Rank = iota
)

const (
	FILE_A File = iota
	FILE_B File = iota
	FILE_C File = iota
	FILE_D File = iota
	FILE_E File = iota
	FILE_F File = iota
	FILE_G File = iota
	FILE_H File = iota
)
