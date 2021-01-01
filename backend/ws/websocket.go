package ws

import (
	"fmt"
	"net/http"

	"github.com/easydaniel/werewolves/backend/game"
	"github.com/gorilla/websocket"
)

type WebSocket struct {
	ws     *websocket.Upgrader
	games  map[string]*game.Game
	player map[string]*game.Player
	action map[string]WSFunction
}

func NewWebSocket() *WebSocket {
	w := new(WebSocket)
	w.ws = &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	w.games = make(map[string]*game.Game)
	w.player = make(map[string]*game.Player)
	w.action = make(map[string]WSFunction)
	w.action["ping"] = Ping

	return w
}

type RecvMessage struct {
	Action string
	Data   interface{}
}

func (w *WebSocket) Handler(writer http.ResponseWriter, req *http.Request) {
	ws, err := w.ws.Upgrade(writer, req, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		var data RecvMessage
		err := ws.ReadJSON(&data)
		if err != nil {
			ws.WriteJSON(map[string]interface{}{
				"error": err.Error(),
			})
			break
		}
		if _, ok := w.action[data.Action]; !ok {
			ws.WriteJSON(map[string]interface{}{
				"error": fmt.Errorf("No this action").Error(),
			})
			break
		}
		res, err := w.action[data.Action](data)
		if err != nil {
			ws.WriteJSON(map[string]interface{}{
				"error": err.Error(),
			})
			break
		}
		ws.WriteJSON(res)
	}
}
