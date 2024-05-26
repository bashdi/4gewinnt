package main

import (
	viergewinnt "4gewinnt/vierGewinnt"
	"fmt"
)

func main() {

	vg := viergewinnt.NewVierGewinnt(25, 10)
	vg.AddPlayer(viergewinnt.NewVgPlayerConsoleInput("Spieler1"))
	vg.AddPlayer(viergewinnt.NewVgPlayerComputer("Computer2"))
	vg.AddPlayer(viergewinnt.NewVgPlayerComputer("Computer3"))
	vg.AddPlayer(viergewinnt.NewVgPlayerComputer("Computer4"))

	vg.SetAnnouncePlayer(func(s string) {
		fmt.Printf("%s ist dran\n", s)
	})

	vg.SetAnnounceDraw(func() {
		fmt.Println("Unentschieden! Es konnte kein Gewinner ermittelt werden!")
	})

	vg.SetAnnounceWinner(func(s string) {
		fmt.Printf("%s hat gewonnen!\n", s)
	})

	vg.SetDrawBoard(func(board [][]int) {
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
	})

	vg.StartGame()
}
