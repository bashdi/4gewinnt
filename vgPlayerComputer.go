package main

import (
	"math/rand"
)

type VgPlayerComputer struct {
	name string
}

func NewVgPlayerComputer(name string) *VgPlayerComputer {
	return &VgPlayerComputer{name: name}
}

func (vgpc *VgPlayerComputer) GetPlayerName() string {
	return vgpc.name
}

func (vgpc *VgPlayerComputer) DoTurn(board [][]int) int {
	column := rand.Intn(len(board[0]))
	return column
}
