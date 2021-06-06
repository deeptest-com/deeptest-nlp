package service

import "github.com/kataras/neffos"

var (
	WsConn *neffos.Conn
)

type WebSocketService struct {
}

func NewWebSocketService() *WebSocketService {
	return &WebSocketService{}
}
