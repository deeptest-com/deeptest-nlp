package net

import (
	"context"
	"fmt"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/log"
	_const "github.com/utlai/utl/internal/pkg/const"
	_domain "github.com/utlai/utl/internal/pkg/domain"
	"testing"
)

func TestTcpClient(t *testing.T) {
	url := fmt.Sprintf("tcp@127.0.0.1:%d", _const.RpcPort)
	d := client.NewPeer2PeerDiscovery(url, "")

	xClient := client.NewXClient("arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xClient.Close()

	args := &_domain.Args{
		A: 1,
		B: 2,
	}

	reply := &_domain.Reply{}

	err := xClient.Call(context.Background(), "Add", args, reply)
	if err != nil {
		log.Errorf("failed to call: %v", err)
	}

	log.Infof("%d + %d = %d", args.A, args.B, reply.C)
}
