package main

import (
	viergewinnt "4gewinnt/vierGewinnt"
)

func main() {

	vg := viergewinnt.NewVierGewinnt(&CmdLineRepresentation{}, 25, 10)
	vg.AddPlayer(NewVgPlayerConsoleInput("Spieler1"))
	vg.AddPlayer(viergewinnt.NewVgPlayerComputer("Computer2"))
	vg.AddPlayer(viergewinnt.NewVgPlayerComputer("Computer3"))
	vg.AddPlayer(viergewinnt.NewVgPlayerComputer("Computer4"))

	vg.StartGame()
}
