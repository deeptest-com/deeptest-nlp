package agentConf

import (
	"github.com/utlai/utl/internal/agent/agentModel"
	consts "github.com/utlai/utl/internal/comm/const"
	_commonUtils "github.com/utlai/utl/internal/pkg/libs/common"
	_fileUtils "github.com/utlai/utl/internal/pkg/libs/file"
	_httpUtils "github.com/utlai/utl/internal/pkg/libs/http"
	_i118Utils "github.com/utlai/utl/internal/pkg/libs/i118"
	"os/user"
	"path/filepath"
)

var (
	Inst = agentModel.Config{}
)

func Init() {
	_i118Utils.InitI118(Inst.Language, consts.AppNameAgent)

	ip, mac, hostName := _commonUtils.GetIp()
	Inst.HostName = hostName
	if Inst.Ip == "" {
		Inst.Ip = ip.String()
	}
	if Inst.MacAddress == "" {
		Inst.MacAddress = mac.String()
	}

	usr, _ := user.Current()
	Inst.WorkDir = _fileUtils.AddPathSepIfNeeded(filepath.Join(usr.HomeDir, "utl"))
	Inst.Server = _httpUtils.UpdateUrl(Inst.Server)
}
