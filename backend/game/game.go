package game

type Game struct {
	ID     string
	Board  *Board
	Day    int
	Player []*Player
}

func NewGame(boardname string) (*Game, error) {
	var err error
	game := new(Game)
	game.Board, err = NewBoard(boardname)
	if err != nil {
		return nil, err
	}
	game.Day = 0
	game.Player = make([]*Player, len(game.Board.Characters))
	return game, err
}
