package viergewinnt

type VgPlayer interface {
	GetPlayerName() string
	DoTurn(board [][]int) int
}
