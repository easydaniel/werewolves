package game

type Player struct {
	Name      string     `json:"name"`
	Character *Character `json:"character"`
	Alive     bool       `json:"alive"`
}
