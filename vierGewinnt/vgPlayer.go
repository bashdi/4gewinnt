package viergewinnt

type VgPlayer interface {
	GetPlayerName() string
	DoTurn(gameState VgGameState) int
}
