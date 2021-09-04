package _resUtils

import (
	_commonUtils "github.com/utlai/utl/internal/pkg/libs/common"
	agentRes "github.com/utlai/utl/res/agent"
	"io/ioutil"
)

func ReadRes(path string) (ret []byte, err error) {
	isRelease := _commonUtils.IsRelease()
	if isRelease {
		ret, err = agentRes.Asset(path)
	} else {
		ret, err = ioutil.ReadFile(path)
	}

	return
}
