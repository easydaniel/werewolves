package game

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type BoardCharacter struct {
	Name string `json:"name"`
	Team int    `json:"team"`
}

type BoardFileCharacter struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type BoardFile struct {
	Name       string                `json:"name"`
	HasSheriff bool                  `json:"has_sheriff"`
	Characters []*BoardFileCharacter `json:"characters"`
}

type Board struct {
	Name       string            `json:"name"`
	Characters []*BoardCharacter `json:"characters"`
	HasSheriff bool              `json:"has_sheriff"`
	NightFlow  []string          `json:"night_flow"`
}

func NewBoard(boardname string) (*Board, error) {
	board := new(Board)
	board.Name = boardname
	filebytes, err := ioutil.ReadFile(fmt.Sprintf("./data/boards/%s.json", boardname))
	if err != nil {
		return nil, err
	}
	var boardfile BoardFile
	json.Unmarshal(filebytes, &boardfile)
	board.HasSheriff = boardfile.HasSheriff
	board.Characters = make([]*BoardCharacter, 0)

	for _, character := range boardfile.Characters {
		char, err := NewCharacter(character.Name)
		if err != nil {
			return nil, err
		}
		for i := 0; i < character.Count; i++ {
			board.Characters = append(board.Characters, &BoardCharacter{
				Name: char.Name,
				Team: char.Team,
			})
		}
		if len(char.Hint) > 0 {
			board.NightFlow = append(board.NightFlow, char.Hint)
		}
	}

	return board, nil
}
