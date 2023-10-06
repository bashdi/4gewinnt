package main

import "fmt"

func main() {

	vg := NewVierGewinnt(18, 5)
	vg.AddPlayer(NewVgPlayerConsoleInput("Spieler1"))
	vg.AddPlayer(NewVgPlayerComputer("Computer2"))
	vg.AddPlayer(NewVgPlayerComputer("Computer3"))
	vg.AddPlayer(NewVgPlayerComputer("Computer4"))

	vg.announcePlayer = func(s string) {
		fmt.Printf("%s ist dran\n", s)
	}

	vg.announceDraw = func() {
		fmt.Println("Unentschieden! Es konnte kein Gewinner ermittelt werden!")
	}

	vg.announceWinner = func(s string) {
		fmt.Printf("%s hat gewonnen!\n", s)
	}

	vg.drawBoard = func(board [][]int) {
		for i := range board {
			for _, field := range board[i] {
				fmt.Printf(" # %d", field)
			}
			fmt.Printf(" #\n")
		}
	}

	vg.StartGame()
}
