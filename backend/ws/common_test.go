package ws

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gorilla/websocket"
)

func initTest() (*WebSocket, *httptest.Server, *websocket.Conn, func(), error) {
	handler := NewWebSocket()

	s := httptest.NewServer(http.HandlerFunc(handler.Handler))

	u := "ws" + strings.TrimPrefix(s.URL, "http")

	ws, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return handler, s, ws, func() {
		s.Close()
		ws.Close()
	}, nil
}
