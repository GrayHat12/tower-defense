package logic

import (
	"math/rand/v2"
)

type TowerDefenseGame struct {
	InitialBoard [EXPECTED_DIM * EXPECTED_DIM]BoardPiece
}

func (game *TowerDefenseGame) ProcedurallyGenerateLevel() {
	// Set castle
	starting_col, ending_col := 1, EXPECTED_DIM-2
	selected_starting_col := rand.IntN(ending_col) + starting_col
	game.InitialBoard[selected_starting_col] = CASTLE

	// Procedurally generate Dirt path

}
