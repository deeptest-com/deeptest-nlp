package service

import (
	"encoding/json"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos"
)

var (
	WsConn *neffos.Conn
)

type WebSocketService struct {
}

func NewWebSocketService() *WebSocketService {
	return &WebSocketService{}
}

func (s *WebSocketService) Broadcast(namespace, room, event string, data interface{}) {
	bytes, _ := json.Marshal(data)

	WsConn.Server().Broadcast(nil, websocket.Message{
		Namespace: namespace,
		Room:      room,
		Event:     event,
		Body:      bytes,
	})
}
