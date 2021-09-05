package serverService

import (
	"bytes"
	"fmt"
	gateway "github.com/rpcx-ecosystem/rpcx-gateway"
	"github.com/smallnest/rpcx/codec"
	_domain "github.com/utlai/utl/internal/pkg/domain"
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"github.com/utlai/utl/internal/server/domain"
	"github.com/utlai/utl/internal/server/model"
	"io/ioutil"
	"net/http"
)

type RpcService struct{}

func NewRpcService() *RpcService {
	return &RpcService{}
}

func (s *RpcService) ExecInstruction(resp serverDomain.NluResp, agent model.Agent) (result _domain.RpcResult) {
	if resp.RasaResult == nil {
		_logUtils.Infof("no intent to exec")
	}

	obj := interface{}(resp.RasaResult)
	result = s.Request(agent.Ip, agent.Port, "selenium", "Exec", &obj)

	result.Pass(fmt.Sprintf("success to send rpc build request %#v.", result))
	return
}

func (s *RpcService) Request(ip string, port int, apiPath string, method string, param *interface{}) (rpcResult _domain.RpcResult) {
	cc := &codec.MsgpackCodec{}

	data, _ := cc.Encode(param)
	url := fmt.Sprintf("http://%s:%d/", ip, port)
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		_logUtils.Errorf("Fail to create request: %#v", err)
		return
	}

	// 设置header
	h := req.Header
	h.Set(gateway.XServicePath, apiPath)
	h.Set(gateway.XServiceMethod, method)
	h.Set(gateway.XMessageID, "10000")
	h.Set(gateway.XMessageType, "0")
	h.Set(gateway.XSerializeType, "3")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		_logUtils.Errorf("fail to call: %#v.", err)
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		_logUtils.Errorf("fail to read response: %#v.", err)
	}

	err = cc.Decode(body, &rpcResult)
	if err != nil {
		_logUtils.Errorf("fail to decode reply: %s.", err.Error())
	}

	msg := fmt.Sprintf("agent return %d-%s.", rpcResult.Code, rpcResult.Msg)
	_logUtils.Info(msg)
	return
}
