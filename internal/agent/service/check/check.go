package checkService

import (
	agentConf "github.com/utlai/utl/internal/agent/conf"
	testService "github.com/utlai/utl/internal/agent/service/test"
	vmService "github.com/utlai/utl/internal/agent/service/vm"
)

func Check() {
	if agentConf.IsVmAgent() { // vm
		// is running，register busy
		if testService.CheckTaskRunning() {
			vmService.RegisterVm(true)
			return
		}

		// no task to run, submit free
		if testService.GetTaskSize() == 0 {
			vmService.RegisterVm(false)
			return
		}

		// has task to run，register busy, then run
		build := testService.PeekTask()
		vmService.RegisterVm(true)
		testService.Exec(build)

	}
}
