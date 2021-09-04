package main

import (
	"flag"
	"github.com/fatih/color"
	"github.com/utlai/utl/cmd/agent/agent"
	"github.com/utlai/utl/cmd/agent/router"
	agentConf "github.com/utlai/utl/internal/agent/conf"
	agentUntils "github.com/utlai/utl/internal/agent/utils/common"
	consts "github.com/utlai/utl/internal/comm/const"
	_const "github.com/utlai/utl/internal/pkg/const"
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"os"
	"os/signal"
	"syscall"
)

var (
	help     bool
	flagSet  *flag.FlagSet
	platform string
)

func main() {
	channel := make(chan os.Signal)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-channel
		cleanup()
		os.Exit(0)
	}()

	flagSet = flag.NewFlagSet("utl", flag.ContinueOnError)

	flagSet.StringVar(&agentConf.Inst.Server, "s", "", "")
	flagSet.StringVar(&agentConf.Inst.Ip, "i", "", "")
	flagSet.IntVar(&agentConf.Inst.Port, "p", _const.RpcPort, "")
	flagSet.StringVar(&agentConf.Inst.Language, "l", "zh", "")

	flagSet.BoolVar(&help, "h", false, "")

	if len(os.Args) == 1 {
		os.Args = append(os.Args, "-h")
	}

	_logUtils.Init(consts.AppNameServer)

	switch os.Args[1] {
	case "help", "-h":
		agentUntils.PrintUsage()

	default:
		if err := flagSet.Parse(os.Args[1:]); err == nil {
			agent.Launch(agentRouter.NewRouter())
		}
	}
}

func init() {
	cleanup()
}

func cleanup() {
	color.Unset()
}
