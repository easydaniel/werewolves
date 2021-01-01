package game

import (
	"fmt"
	"math/rand"
	"sync"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type Game struct {
	ID     string
	Board  *Board
	Day    int
	Player []*Player
	Host   *Player
	Guest  map[string]*Player

	lock *sync.Mutex
}

func NewGame(boardname string, host *Player) (*Game, error) {
	var err error
	game := new(Game)
	game.Board, err = NewBoard(boardname)
	if err != nil {
		return nil, err
	}
	game.Day = 0
	game.Player = make([]*Player, len(game.Board.Characters))
	game.Host = host
	game.lock = new(sync.Mutex)
	return game, err
}

func (g *Game) Start() error {
	g.lock.Lock()
	defer g.lock.Unlock()

	shuffleList := []int{}
	for i, player := range g.Player {
		if player == nil {
			return fmt.Errorf("Player %v is empty", i+1)
		}
		shuffleList = append(shuffleList, i)
	}
	g.Day = 1

	rand.Shuffle(len(g.Board.Characters), func(i, j int) {
		shuffleList[i], shuffleList[j] = shuffleList[j], shuffleList[i]
	})
	for i, charID := range shuffleList {
		g.Player[i].Character = g.Board.Characters[charID]
	}
	return nil
}

func (g *Game) SetSeat(player *Player, id int) error {
	g.lock.Lock()
	defer g.lock.Unlock()
	if g.Day != 0 {
		return fmt.Errorf("Game Already Started")
	}
	if _, ok := g.Guest[player.Name]; !ok {
		return fmt.Errorf("Player Not In Room")
	}
	if g.Player[id] != nil {
		return fmt.Errorf("Seat Already Have Player")
	}
	for i, p := range g.Player {
		if player.Name == p.Name {
			g.Player[i] = nil
		}
	}
	g.Player[id] = player
	return nil
}

func (g *Game) ExitSeat(player *Player, id int) error {
	g.lock.Lock()
	defer g.lock.Unlock()
	if g.Day != 0 {
		return fmt.Errorf("Game Already Started")
	}
	for i, p := range g.Player {
		if player.Name == p.Name {
			g.Player[i] = nil
			return nil
		}
	}
	return fmt.Errorf("Player Not On The Seat")
}

func (g *Game) JoinRoom(player *Player) error {
	g.lock.Lock()
	defer g.lock.Unlock()
	if _, ok := g.Guest[player.Name]; ok {
		return fmt.Errorf("Player Already In Room")
	}
	g.Guest[player.Name] = player
	return nil
}

func (g *Game) ExitRoom(player *Player) error {
	g.lock.Lock()
	defer g.lock.Unlock()
	if _, ok := g.Guest[player.Name]; !ok {
		return fmt.Errorf("Player Not In Room")
	}
	g.Guest[player.Name] = nil
	return nil
}
