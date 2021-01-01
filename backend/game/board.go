package game

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

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
	Name       string       `json:"name"`
	Characters []*Character `json:"characters"`
	HasSheriff bool         `json:"has_sheriff"`
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
	board.Characters = make([]*Character, 0)

	for _, character := range boardfile.Characters {
		for i := 0; i < character.Count; i++ {
			char, err := NewCharacter(character.Name)
			if err != nil {
				return nil, err
			}
			board.Characters = append(board.Characters, char)
		}
	}

	return board, nil
}
