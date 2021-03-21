package vmService

import (
	agentConf "github.com/utlai/utl/internal/agent/conf"
	_const "github.com/utlai/utl/internal/pkg/const"
	_domain "github.com/utlai/utl/internal/pkg/domain"
	_httpUtils "github.com/utlai/utl/internal/pkg/libs/http"
	_i118Utils "github.com/utlai/utl/internal/pkg/libs/i118"
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
)

func RegisterVm(isBusy bool) {
	vm := _domain.Vm{HostName: agentConf.Inst.HostName, WorkDir: agentConf.Inst.WorkDir,
		PublicIp: agentConf.Inst.Ip, PublicPort: agentConf.Inst.Port}

	if isBusy {
		vm.Status = _const.VmBusy
	} else {
		vm.Status = _const.VmActive
	}

	url := _httpUtils.GenUrl(agentConf.Inst.Server, "vms/register")
	resp, ok := _httpUtils.Post(url, vm)

	if ok {
		_logUtils.Info(_i118Utils.I118Prt.Sprintf("success_to_register", agentConf.Inst.Server))
	} else {
		_logUtils.Info(_i118Utils.I118Prt.Sprintf("fail_to_register", agentConf.Inst.Server, resp))
	}
}
