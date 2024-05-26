package viergewinnt

type VgRepresentation interface {
	DrawBoard([][]int)

	AnnouncePlayersTurn(string)
	AnnounceWinner(string)
	AnnounceDraw()
}
