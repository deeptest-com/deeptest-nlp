package manageService

import (
	"fmt"
	managerConf "github.com/utlai/utl/cmd/manager/utils/conf"
	managerVari "github.com/utlai/utl/cmd/manager/utils/vari"
	_const "github.com/utlai/utl/internal/pkg/const"
	_commonUtils "github.com/utlai/utl/internal/pkg/libs/common"
	_fileUtils "github.com/utlai/utl/internal/pkg/libs/file"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

func StartApp(app managerConf.Client) (err error) {
	appDir := managerVari.WorkDir + app.Name + _const.PthSep

	newExePath := path.Join(appDir, app.Version, app.Name, app.Name)
	if _commonUtils.IsWin() {
		newExePath += ".exe"
	}

	execDir := _fileUtils.AbsolutePath(filepath.Dir(newExePath))

	tmpl := ""
	cmdStr := ""
	var cmd *exec.Cmd
	if _commonUtils.IsWin() {
		tmpl = `start cmd /c %s %s ^1^> %snohup.%s.log ^2^>^&^1`
		cmdStr = fmt.Sprintf(tmpl, newExePath, app.Params, managerVari.WorkDir, app.Name)

		cmd = exec.Command("cmd", "/C", cmdStr)

	} else {
		cmd = exec.Command("nohup", newExePath, app.Params)

		log := filepath.Join(managerVari.WorkDir, "nohup."+app.Name+".log")
		f, _ := os.Create(log)

		cmd.Stdout = f
		cmd.Stderr = f
	}

	cmd.Dir = execDir
	err = cmd.Start()

	return
}
