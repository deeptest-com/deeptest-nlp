package _shellUtils

import (
	"bufio"
	"bytes"
	"fmt"
	_commonUtils "github.com/utlai/utl/internal/pkg/libs/common"
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	_stringUtils "github.com/utlai/utl/internal/pkg/libs/string"
	"io"
	"os/exec"
	"regexp"
	"strings"
)

func ExeSysCmd(cmdStr string) (string, error) {
	var cmd *exec.Cmd
	if _commonUtils.IsWin() {
		cmd = exec.Command("cmd", "/C", cmdStr)
	} else {
		cmd = exec.Command("/bin/bash", "-c", cmdStr)
	}

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	output := out.String()

	return output, err
}

func ExeShell(cmdStr string, dir string) (ret string, err error, pid int) {
	var cmd *exec.Cmd
	if _commonUtils.IsWin() {
		cmd = exec.Command("cmd", "/C", cmdStr)
	} else {
		cmd = exec.Command("/bin/bash", "-c", cmdStr)
	}
	if dir != "" {
		cmd.Dir = dir
	}

	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		_logUtils.Error(fmt.Sprintf("fail to exec command `%s` in dir `%s`, error `%#v`.", cmdStr, cmd.Dir, err))
	}

	pid = cmd.Process.Pid
	ret = _stringUtils.TrimAll(out.String())
	return
}

func ExeShellWithOutput(cmdStr string, dir string) (output []string, err error, pid int) {
	var cmd *exec.Cmd
	if _commonUtils.IsWin() {
		cmd = exec.Command("cmd", "/C", cmdStr)
	} else {
		cmd = exec.Command("/bin/bash", "-c", cmdStr)
	}

	if dir != "" {
		cmd.Dir = dir
	}

	stdout, err1 := cmd.StdoutPipe()
	if err1 != nil {
		err = err1
		_logUtils.Error(err.Error())
		return
	}

	cmd.Start()

	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		_logUtils.Info(strings.TrimRight(line, "\n"))
		output = append(output, line)
	}

	cmd.Wait()

	pid = cmd.Process.Pid
	return
}

func GetProcess(app string) (string, error) {
	var cmd *exec.Cmd

	tmpl := ""
	cmdStr := ""
	if _commonUtils.IsWin() {
		tmpl = `tasklist`
		cmdStr = fmt.Sprintf(tmpl)

		cmd = exec.Command("cmd", "/C", cmdStr)
	} else {
		tmpl = `ps -ef | grep "%s" | grep -v "grep" | awk '{print $2}'`
		cmdStr = fmt.Sprintf(tmpl, app)

		cmd = exec.Command("/bin/bash", "-c", cmdStr)
	}

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	output := ""
	if _commonUtils.IsWin() {
		arr := strings.Split(out.String(), "\n")
		for _, line := range arr {
			if strings.Index(line, app+".exe") > -1 {
				arr2 := regexp.MustCompile(`\s+`).Split(line, -1)
				output = arr2[1]
				break
			}
		}
	} else {
		output = out.String()
	}

	return output, err
}

func KillProcess(app string) (string, error) {
	var cmd *exec.Cmd

	tmpl := ""
	cmdStr := ""
	if _commonUtils.IsWin() {
		// tasklist | findstr ztf.exe
		tmpl = `taskkill.exe /f /im %s.exe`
		cmdStr = fmt.Sprintf(tmpl, app)

		cmd = exec.Command("cmd", "/C", cmdStr)
	} else {
		tmpl = `ps -ef | grep '%s' | grep -v "grep" | awk '{print $2}' | xargs kill -9`
		cmdStr = fmt.Sprintf(tmpl, app)

		cmd = exec.Command("/bin/bash", "-c", cmdStr)
	}

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	output := out.String()

	return output, err
}

func KillProcessById(pid int) {
	cmdStr := fmt.Sprintf("kill -9 %d", pid)
	ExeShell(cmdStr, "")
}
