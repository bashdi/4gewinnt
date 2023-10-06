package main

type VgPlayer interface {
	GetPlayerName() string
	DoTurn(board [][]int) int
}
