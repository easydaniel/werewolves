package game

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Character struct {
	Name string `json:"name"`
	Team int    `json:"team"`
	Hint string `json:"hint"`
}

func NewCharacter(charactername string) (*Character, error) {
	character := new(Character)
	filebytes, err := ioutil.ReadFile(fmt.Sprintf("./data/characters/%s.json", charactername))
	if err != nil {
		return nil, err
	}
	json.Unmarshal(filebytes, character)
	return character, nil
}
