package logic

const EXPECTED_DIM = 14

type BoardPiece int

const (
	F_GRASS BoardPiece = iota
	U_GRASS
	R_GRASS
	D_GRASS
	L_GRASS
	DIRT
	TOWER
	CASTLE
)
