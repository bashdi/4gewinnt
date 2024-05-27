package viergewinnt

type VgGameState struct {
	currentPlayer int
	board         [][]int
}

func (gs *VgGameState) CurrentPlayer() int {
	return gs.currentPlayer
}

func (gs *VgGameState) Board() [][]int {
	return gs.board
}
