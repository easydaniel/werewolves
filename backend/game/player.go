package game

type Player struct {
	Name      string          `json:"name"`
	Character *BoardCharacter `json:"character"`
	Alive     bool            `json:"alive"`
}
