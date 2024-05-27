package viergewinnt

import (
	"math/rand"
)

// Simple Implementation einer Spieler-KI
type VgPlayerComputer struct {
	name string
}

func NewVgPlayerComputer(name string) *VgPlayerComputer {
	return &VgPlayerComputer{name: name}
}

func (vgpc *VgPlayerComputer) GetPlayerName() string {
	return vgpc.name
}

func (vgpc *VgPlayerComputer) DoTurn(gameState VgGameState) int {
	column := rand.Intn(len(gameState.Board()[0]))
	return column
}
