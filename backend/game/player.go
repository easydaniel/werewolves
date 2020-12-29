package game

import "github.com/gorilla/websocket"

type Player struct {
	Name      string
	Password  string
	Character *Character
	Connect   *websocket.Conn
}
