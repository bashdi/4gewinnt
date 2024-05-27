package main

import (
	viergewinnt "4gewinnt/vierGewinnt"
	"fmt"
)

type CmdLineRepresentation struct {
}

func (clr *CmdLineRepresentation) DrawBoard(board [][]int) {
	for i := range board {
		for _, field := range board[i] {
			if field == 0 {
				fmt.Printf(" #  ")
			} else {
				fmt.Printf(" # %d", field)
			}

		}
		fmt.Printf(" #\n")
	}
}

func (clr *CmdLineRepresentation) AnnouncePlayersTurn(s string) {
	fmt.Printf("%s ist dran\n", s)
}

func (clr *CmdLineRepresentation) AnnounceWinner(s string) {
	fmt.Printf("%s hat gewonnen!\n", s)
}

func (clr *CmdLineRepresentation) AnnounceDraw() {
	fmt.Println("Unentschieden! Es konnte kein Gewinner ermittelt werden!")
}

type VgPlayerCmdLine struct {
	name string
}

func NewVgPlayerConsoleInput(name string) *VgPlayerCmdLine {
	return &VgPlayerCmdLine{name: name}
}

func (vgpc *VgPlayerCmdLine) GetPlayerName() string {
	return vgpc.name
}

func (vgpc *VgPlayerCmdLine) DoTurn(gameState viergewinnt.VgGameState) int {
	fmt.Println("Enter column:")
	var column int
	_, err := fmt.Scanf("%d", &column)
	if err != nil {
		return -1
	}
	return column
}
