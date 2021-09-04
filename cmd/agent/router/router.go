package agentRouter

import (
	"fmt"
	"github.com/smallnest/rpcx/server"
	agentHandler "github.com/utlai/utl/cmd/agent/router/handler"
	agentConf "github.com/utlai/utl/internal/agent/conf"
	_i118Utils "github.com/utlai/utl/internal/pkg/libs/i118"
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
)

type Router struct {
	ArithCtrl    *agentHandler.ArithCtrl    `inject:""`
	SeleniumCtrl *agentHandler.SeleniumCtrl `inject:""`
}

func NewRouter() *Router {
	router := &Router{}
	return router
}

func (r *Router) App() {
	addr := fmt.Sprintf("0.0.0.0:%d", agentConf.Inst.Port)
	srv := server.NewServer()

	srv.RegisterName("arith", r.ArithCtrl, "")
	srv.RegisterName("selenium", r.SeleniumCtrl, "")

	_logUtils.Info(_i118Utils.Sprintf("start_server", addr))
	err := srv.Serve("tcp", addr)
	if err != nil {
		_logUtils.Infof(_i118Utils.Sprintf("fail_to_start_server", addr, err.Error()))
	}
}
