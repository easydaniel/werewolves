package ws

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	_, _, ws, closeFunc, err := initTest()
	assert.Nil(t, err)
	defer closeFunc()

	ws.WriteJSON(map[string]interface{}{
		"action": "ping",
	})

	var res interface{}
	expected := map[string]interface{}{
		"message": "pong",
	}

	ws.ReadJSON(&res)
	assert.Equal(t, res, expected)
}
