package handler

import (
	"fmt"
	"github.com/kataras/iris/v12/websocket"
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"github.com/utlai/utl/internal/server/service"
	serverConst "github.com/utlai/utl/internal/server/utils/const"
)

type WsCtrl struct {
	BaseCtrl
	Namespace         string
	*websocket.NSConn `stateless:"true"`

	WebSocketService *service.WebSocketService `inject:""`
}

func NewWsCtrl() *WsCtrl {
	return &WsCtrl{Namespace: serverConst.WsNamespace}
}

func (c *WsCtrl) OnNamespaceConnected(msg websocket.Message) error {
	service.WsConn = c.Conn

	_logUtils.Infof("websocket client connected %s", c.Conn.ID())
	c.Conn.Server().Broadcast(nil, websocket.Message{
		Namespace: msg.Namespace,
		Event:     "OnVisit",
		Body:      []byte(fmt.Sprintf("%s", "server: connected to websocket")),
	})
	return nil
}

// OnNamespaceDisconnect This will call the "OnVisit" event on all clients, except the current one,
// it can't because it's left but for any case use this type of design
func (c *WsCtrl) OnNamespaceDisconnect(msg websocket.Message) error {
	_logUtils.Infof("%s disconnected", c.Conn.ID())
	c.Conn.Server().Broadcast(nil, websocket.Message{
		Namespace: msg.Namespace,
		Event:     "OnVisit",
		Body:      []byte(fmt.Sprintf("%s", "server: disconnected from websocket")),
	})
	return nil
}

// OnChat This will call the "OnVisit" event on all clients, including the current one, with the 'newCount' variable.
func (c *WsCtrl) OnChat(msg websocket.Message) (err error) {
	ctx := websocket.GetContext(c.Conn)

	str := ctx.RemoteAddr()
	_logUtils.Info(str + ", " + string(msg.Body))

	c.Conn.Server().Broadcast(nil, websocket.Message{
		Namespace: msg.Namespace,
		Room:      msg.Room,
		Event:     msg.Event,
		Body:      []byte(fmt.Sprintf(" response %s", "abc")),
	})

	return
}
