package agent

import (
	"github.com/facebookgo/inject"
	"github.com/sirupsen/logrus"
	agentRouter "github.com/utlai/utl/cmd/agent/router"
	agentConf "github.com/utlai/utl/internal/agent/conf"
	agentCron "github.com/utlai/utl/internal/agent/cron"
	consts "github.com/utlai/utl/internal/comm/const"
	_db "github.com/utlai/utl/internal/pkg/db"
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
)

func Launch(router *agentRouter.Router) {
	_logUtils.Init(consts.AppNameAgent)

	agentConf.Init()

	_db.InitDB("agent")

	injectObj(router)
	router.App()
}

func injectObj(router *agentRouter.Router) {
	// inject
	var g inject.Graph
	g.Logger = logrus.StandardLogger()

	err := g.Provide(
		// db
		&inject.Object{Value: _db.GetInst().DB()},

		// cron
		&inject.Object{Value: agentCron.NewAgentCron()},

		// router, its controllers etc.
		&inject.Object{Value: router},
	)

	if err != nil {
		logrus.Fatalf("provide usecase objects to the Graph: %v", err)
	}

	err = g.Populate()
	if err != nil {
		logrus.Fatalf("populate the incomplete Objects: %v", err)
	}
}
