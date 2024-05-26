package viergewinnt

type VierGewinnt struct {
	xSize         int
	ySize         int
	board         [][]int
	players       []VgPlayer
	currentPlayer int

	visualAction VgRepresentation
}

func NewVierGewinnt(visualAction VgRepresentation, xSize, ySize int) *VierGewinnt {
	vg := VierGewinnt{visualAction: visualAction, xSize: xSize, ySize: ySize}
	vg.createBoard()

	return &vg
}

// Spielfeld erstellen
func (vg *VierGewinnt) createBoard() {
	//cordSystem aufbauen
	vg.board = make([][]int, vg.ySize)

	for i := 0; i < vg.ySize; i++ {
		row := make([]int, vg.xSize)
		vg.board[i] = row
	}
}

// Spielfeld zurücksetzen
func (vg *VierGewinnt) ClearBoard() {
	for i := 0; i < len(vg.board); i++ {
		clear(vg.board[i])
	}
}

// Spieler hinzufügen
func (vg *VierGewinnt) AddPlayer(player VgPlayer) {
	vg.players = append(vg.players, player)
}

// Aktueller Spieler setzt seine Münze
func (vg *VierGewinnt) addCoin(column int) bool {
	//Überprüfen ob es die Spalte gibt
	if column >= vg.xSize || column < 0 {
		return false
	}

	//Gravitation der Spalte
	for i := vg.ySize - 1; i >= 0; i-- {
		if vg.board[i][column] == 0 {
			vg.board[i][column] = vg.currentPlayer
			return true
		}
	}
	return false
}

// Prüft ob es einen Gewinner gibt
func (vg *VierGewinnt) isThereAWinner() bool {
	//Max cords
	maxY := len(vg.board) - 1
	maxX := len(vg.board[0]) - 1

	//überprüfungs reichweite
	checkRange := 3

	for y, rows := range vg.board {
		for x, field := range rows {
			//Bei 0 kann eine Überprüfung ignoriert werden
			if field == 0 {
				continue
			}

			//horizontal nach rechts prüfen
			for i := 1; i <= checkRange; i++ {
				if x+i > maxX {
					break
				}
				if vg.board[y][x+i] != field {
					break
				}
				if i >= checkRange {
					return true
				}
			}

			//vertikal
			for i := 1; i <= checkRange; i++ {
				if y+i > maxY {
					break
				}
				if vg.board[y+i][x] != field {
					break
				}
				if i >= checkRange {
					return true
				}
			}

			//diagonal 1 nach rechts prüfen
			for i := 1; i <= checkRange; i++ {
				if y+i > maxY || x+i > maxX {
					break
				}
				if vg.board[y+i][x+i] != field {
					break
				}
				if i >= checkRange {
					return true
				}
			}

			//diagonal 2 nach rechts prüfen
			for i := 1; i <= checkRange; i++ {
				if y-i < 0 || x+i > maxX {
					break
				}
				if vg.board[y-i][x+i] != field {
					break
				}
				if i >= checkRange {
					return true
				}
			}

		}
	}
	return false
}

// Prüft ob alle Felder schon belegt sind und ein Unentschieden vorliegt
func (vg *VierGewinnt) isThereADraw() bool {
	for _, row := range vg.board {
		for _, field := range row {
			if field == 0 {
				return false
			}
		}
	}
	return true
}

// zu nächstem Spieler wechseln
func (vg *VierGewinnt) nextPlayer() {
	vg.currentPlayer++
	if len(vg.players) < vg.currentPlayer {
		vg.currentPlayer = 1
	}
}

// Startet Gameloop
func (vg *VierGewinnt) StartGame() {
	//init board und erstes zeichnen
	vg.createBoard()
	vg.visualAction.DrawBoard(vg.board)
	vg.nextPlayer()

	//gameloop
	for {
		vgPlayer := vg.players[vg.currentPlayer-1]
		vg.visualAction.AnnouncePlayersTurn(vgPlayer.GetPlayerName())
		//Spieler eingabe, wird solange gemacht bis eingabe korrekt
		for {
			column := vgPlayer.DoTurn(vg.board)

			if vg.addCoin(column) {
				break
			}
		}

		vg.visualAction.DrawBoard(vg.board)

		//Prüfen ob jemand gewonnen hat
		if vg.isThereAWinner() {
			vg.visualAction.AnnounceWinner(vgPlayer.GetPlayerName())
			break
		}

		//Prüfen ob es ein Unentschieden ist
		if vg.isThereADraw() {
			vg.visualAction.AnnounceDraw()
			break
		}

		vg.nextPlayer()
	}
}
