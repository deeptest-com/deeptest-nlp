package middlewareUtils

import (
	"fmt"
	_fileUtils "github.com/utlai/utl/internal/pkg/libs/file"
	"github.com/utlai/utl/internal/pkg/utils"
	"github.com/utlai/utl/internal/server/db"
	"github.com/casbin/casbin/v2"
	"github.com/sirupsen/logrus"
	"path/filepath"
)

func NewEnforcer() *casbin.Enforcer {
	adapter, err := NewAdapterByDB(db.GetInst().DB())
	if err != nil {
		logrus.Println(fmt.Sprintf("NewAdapter 错误: %v", err))
	}

	exeDir := _utils.GetExeDir()
	pth := filepath.Join(exeDir, "rbac_model.conf")
	if !_fileUtils.FileExist(pth) { // debug mode
		pth = filepath.Join(exeDir, "cmd", "server", "rbac_model.conf")
	}

	enforcer, err := casbin.NewEnforcer(pth, adapter)
	if err != nil {
		logrus.Println(fmt.Sprintf("NewEnforcer 错误: %v", err))
	}

	_ = enforcer.LoadPolicy()

	return enforcer
}
