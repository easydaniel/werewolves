package game

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

type BoardFileCharacter struct {
	Name       string `json:"name"`
	Team       int    `json:"team"`
	NightOrder int    `json:"night_order"`
	Hint       string `json:"hint"`
	Count      int    `json:"count"`
}

type BoardFile struct {
	Name       string                `json:"name"`
	HasSheriff bool                  `json:"has_sheriff"`
	Characters []*BoardFileCharacter `json:"characters"`
}

type Board struct {
	Name       string
	Characters []*Character
	HasSheriff bool
}

func NewBoard(boardname string) (*Board, error) {
	board := new(Board)
	board.Name = boardname
	filebytes, err := ioutil.ReadFile(filepath.Join("./board", boardname))
	if err != nil {
		return nil, err
	}
	var boardfile BoardFile
	json.Unmarshal(filebytes, &boardfile)
	board.HasSheriff = boardfile.HasSheriff
	board.Characters = make([]*Character, 0)

	for _, character := range boardfile.Characters {
		for i := 0; i < character.Count; i++ {
			board.Characters = append(board.Characters, &Character{
				Name:       character.Name,
				Team:       character.Team,
				NightOrder: character.NightOrder,
				Hint:       character.Hint,
			})
		}
	}

	return board, nil
}
