package agentConf

import (
	"github.com/utlai/utl/internal/agent/agentModel"
	agentConst "github.com/utlai/utl/internal/agent/utils/const"
	_const "github.com/utlai/utl/internal/pkg/const"
	_commonUtils "github.com/utlai/utl/internal/pkg/libs/common"
	_fileUtils "github.com/utlai/utl/internal/pkg/libs/file"
	_httpUtils "github.com/utlai/utl/internal/pkg/libs/http"
	_i118Utils "github.com/utlai/utl/internal/pkg/libs/i118"
	"os/user"
	"path"
)

var (
	Inst = agentModel.Config{}
)

func Init() {
	_i118Utils.InitI118(Inst.Language, agentConst.AppName)

	ip, _, hostName := _commonUtils.GetIp()
	Inst.HostName = hostName
	Inst.Ip = ip.String()
	Inst.Port = _const.RpcPort

	usr, _ := user.Current()
	Inst.WorkDir = _fileUtils.AddPathSepIfNeeded(path.Join(usr.HomeDir, "utl"))
	Inst.Server = _httpUtils.UpdateUrl(Inst.Server)
}

func IsVmAgent() bool {
	return Inst.Platform == _const.Vm
}
func IsDeviceAgent() bool {
	return IsIosAgent() || IsAndroidAgent()
}
func IsAndroidAgent() bool {
	return Inst.Platform == _const.Android
}

func IsIosAgent() bool {
	return Inst.Platform == _const.Ios
}
