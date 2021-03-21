package deviceCommon

import (
	agentConf "github.com/utlai/utl/internal/agent/conf"
	_domain "github.com/utlai/utl/internal/pkg/domain"
	"github.com/jinzhu/copier"
)

func SpecToDevInsts(specs []_domain.DeviceSpec) []_domain.DeviceInst {
	insts := make([]_domain.DeviceInst, 0)

	for _, spec := range specs {
		inst := _domain.DeviceInst{}
		copier.Copy(&inst, spec)

		inst.ComputerIp = agentConf.Inst.Ip
		inst.ComputerPort = agentConf.Inst.Port

		insts = append(insts, inst)
	}

	return insts
}
