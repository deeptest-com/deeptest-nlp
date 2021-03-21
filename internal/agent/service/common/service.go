package commonService

import (
	agentConf "github.com/utlai/utl/internal/agent/conf"
	_domain "github.com/utlai/utl/internal/pkg/domain"
	_fileUtils "github.com/utlai/utl/internal/pkg/libs/file"
	uuid "github.com/satori/go.uuid"
	"os"
)

func SetBuildWorkDir(build *_domain.BuildTo) {
	build.WorkDir = agentConf.Inst.WorkDir + uuid.NewV4().String() + string(os.PathSeparator)
	_fileUtils.MkDirIfNeeded(build.WorkDir)
}
