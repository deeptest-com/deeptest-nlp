package handler

import (
	"fmt"
	"github.com/kataras/iris/v12/websocket"
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"github.com/utlai/utl/internal/server/service"
	serverConst "github.com/utlai/utl/internal/server/utils/const"
)

type WsCtrl struct {
	CommCtrl
	Namespace         string
	*websocket.NSConn `stateless:"true"`

	WebSocketService *serverService.WebSocketService `inject:""`
}

func NewWsCtrl() *WsCtrl {
	return &WsCtrl{Namespace: serverConst.WsNamespace}
}

func (c *WsCtrl) OnNamespaceConnected(msg websocket.Message) error {
	c.WebSocketService.SetConn(c.Conn)

	_logUtils.Infof("websocket client connected %s", c.Conn.ID())

	data := map[string]string{"msg": "server: connected to websocket"}
	c.WebSocketService.Broadcast(msg.Namespace, "", "OnVisit", data)
	return nil
}

// OnNamespaceDisconnect This will call the "OnVisit" event on all clients, except the current one,
// it can't because it's left but for any case use this type of design
func (c *WsCtrl) OnNamespaceDisconnect(msg websocket.Message) error {
	_logUtils.Infof("%s disconnected", c.Conn.ID())

	data := map[string]string{"msg": "server: disconnected to websocket"}
	c.WebSocketService.Broadcast(msg.Namespace, "", "OnVisit", data)
	return nil
}

// OnChat This will call the "OnVisit" event on all clients, including the current one, with the 'newCount' variable.
func (c *WsCtrl) OnChat(msg websocket.Message) (err error) {
	ctx := websocket.GetContext(c.Conn)

	str := ctx.RemoteAddr()
	_logUtils.Info(str + ", " + string(msg.Body))

	data := map[string]string{"msg": fmt.Sprintf("response %s", "abc")}
	c.WebSocketService.Broadcast(msg.Namespace, msg.Room, msg.Event, data)

	return
}
