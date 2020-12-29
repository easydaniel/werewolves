package ws

func Ping(data interface{}) (interface{}, error) {
	return map[string]interface{}{
		"message": "pong",
	}, nil
}
