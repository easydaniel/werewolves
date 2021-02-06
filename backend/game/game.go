package game

import (
	"fmt"
	"math/rand"
	"sync"
)

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

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
	Host   *Member
	Member map[string]*Member

	lock *sync.Mutex
}

func NewGame(boardname string, host *Member) (*Game, error) {
	var err error
	game := new(Game)
	game.Board, err = NewBoard(boardname)
	if err != nil {
		return nil, err
	}
	game.ID = RandStringRunes(6)
	game.Day = 0
	game.Player = make([]*Player, len(game.Board.Characters))
	game.Member = make(map[string]*Member)
	game.Member[host.Name] = host
	game.Host = host
	game.lock = new(sync.Mutex)
	return game, err
}

// func (g *Game) Start() error {
// 	g.lock.Lock()
// 	defer g.lock.Unlock()

// 	shuffleList := []int{}
// 	for i, player := range g.Player {
// 		if player == nil {
// 			return fmt.Errorf("Player %v is empty", i+1)
// 		}
// 		shuffleList = append(shuffleList, i)
// 	}
// 	g.Day = 1

// 	rand.Shuffle(len(g.Board.Characters), func(i, j int) {
// 		shuffleList[i], shuffleList[j] = shuffleList[j], shuffleList[i]
// 	})
// 	for i, charID := range shuffleList {
// 		g.Player[i].Character = g.Board.Characters[charID]
// 		g.Player[i].Alive = true
// 	}
// 	return nil
// }

func (g *Game) SetSeat(member *Member, id int) error {
	g.lock.Lock()
	defer g.lock.Unlock()
	if g.Day != 0 {
		return fmt.Errorf("Game Already Started")
	}
	if id >= len(g.Player) || id < 0 {
		return fmt.Errorf("Seat ID Should Between 0 and %v", len(g.Player)-1)
	}
	if _, ok := g.Member[member.Name]; !ok {
		return fmt.Errorf("Player Not In Room")
	}
	// Host Cannot set seat
	if member.Name == g.Host.Name {
		return fmt.Errorf("Host Cannot Enter Seat")
	}
	if g.Player[id] != nil {
		return fmt.Errorf("Seat Already Have Player")
	}
	for i, player := range g.Player {
		if player != nil && player.Name == member.Name {
			g.Player[i] = nil
		}
	}
	g.Player[id] = &Player{
		Name: member.Name,
	}
	return nil
}

func (g *Game) ExitSeat(id int) error {
	g.lock.Lock()
	defer g.lock.Unlock()
	if g.Day != 0 {
		return fmt.Errorf("Game Already Started")
	}
	if id >= len(g.Player) || id < 0 {
		return fmt.Errorf("Seat ID Should Between 0 and %v", len(g.Player)-1)
	}
	if g.Player[id] == nil {
		return fmt.Errorf("Player Not On The Seat")
	}
	g.Player[id] = nil
	return nil
}

func (g *Game) JoinRoom(member *Member) error {
	g.lock.Lock()
	defer g.lock.Unlock()
	if _, ok := g.Member[member.Name]; ok {
		return fmt.Errorf("Player Already In Room")
	}
	g.Member[member.Name] = member
	return nil
}

func (g *Game) ExitRoom(member *Member) error {
	g.lock.Lock()
	defer g.lock.Unlock()
	if _, ok := g.Member[member.Name]; !ok {
		return fmt.Errorf("Player Not In Room")
	}
	for i, player := range g.Player {
		if player != nil && player.Name == member.Name {
			g.Player[i] = nil
		}
	}
	delete(g.Member, member.Name)
	return nil
}

func (g *Game) ChangeHost(member *Member) error {
	g.lock.Lock()
	defer g.lock.Unlock()
	if _, ok := g.Member[member.Name]; !ok {
		return fmt.Errorf("Player Not In Room")
	}
	for _, player := range g.Player {
		if player != nil && player.Name == member.Name {
			return fmt.Errorf("Cannot Set Player As Host")
		}
	}
	g.Host = member
	return nil
}
