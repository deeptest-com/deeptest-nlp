package main

import (
	"github.com/utlai/utl/cmd/server/server"
	consts "github.com/utlai/utl/internal/comm/const"
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
)

func main() {
	_logUtils.Init(consts.AppNameServer)

	server.Launch()
}
