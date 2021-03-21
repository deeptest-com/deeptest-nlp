package manageService

import (
	managerConf "github.com/utlai/utl/cmd/manager/utils/conf"
	_shellUtils "github.com/utlai/utl/internal/pkg/libs/shell"
	"strings"
)

func CheckStatus(app managerConf.Client) {
	output, _ := _shellUtils.GetProcess(app.Name)
	output = strings.TrimSpace(output)

	if output != "" {
		return
	}

	StartApp(app)
}
