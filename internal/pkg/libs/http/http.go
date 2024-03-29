package _httpUtils

import (
	"encoding/json"
	"fmt"
	"github.com/utlai/utl/internal/comm/domain"
	_const "github.com/utlai/utl/internal/pkg/const"
	_domain "github.com/utlai/utl/internal/pkg/domain"
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	_vari "github.com/utlai/utl/internal/pkg/vari"
	"io/ioutil"
	"net/http"
	"strings"
)

func Get(url string) (interface{}, bool) {
	return GetObj(url, "farm")
}

func GetObj(url string, requestTo string) (interface{}, bool) {
	client := &http.Client{}

	if _vari.Verbose {
		_logUtils.Info(url)
	}

	req, reqErr := http.NewRequest("GET", url, nil)
	if reqErr != nil {
		_logUtils.Error(reqErr.Error())
		return nil, false
	}

	resp, respErr := client.Do(req)

	if respErr != nil {
		_logUtils.Error(respErr.Error())
		return nil, false
	}

	bodyStr, _ := ioutil.ReadAll(resp.Body)
	if _vari.Verbose {
		_logUtils.PrintUnicode(bodyStr)
	}
	defer resp.Body.Close()

	if requestTo == "farm" {
		var bodyJson _domain.RpcResult
		jsonErr := json.Unmarshal(bodyStr, &bodyJson)
		if jsonErr != nil {
			if strings.Index(string(bodyStr), "<html>") > -1 {
				_logUtils.Error("server return a html")
				return nil, false
			} else {
				_logUtils.Error(jsonErr.Error())
				return nil, false
			}
		}
		code := bodyJson.Code
		return bodyJson.Payload, code == _const.ResultSuccess
	} else {
		var bodyJson map[string]interface{}
		jsonErr := json.Unmarshal(bodyStr, &bodyJson)
		if jsonErr != nil {
			_logUtils.Error(jsonErr.Error())
			return nil, false
		} else {
			return bodyJson, true
		}
	}
}

func Post(url string, params interface{}) (interface{}, bool) {
	if _vari.Verbose {
		_logUtils.Info(url)
	}
	client := &http.Client{}

	paramStr, err := json.Marshal(params)
	if err != nil {
		_logUtils.Error(err.Error())
		return nil, false
	}

	req, reqErr := http.NewRequest("POST", url, strings.NewReader(string(paramStr)))
	if reqErr != nil {
		_logUtils.Error(reqErr.Error())
		return nil, false
	}

	req.Header.Set("Content-Type", "application/json")

	resp, respErr := client.Do(req)
	if respErr != nil {
		_logUtils.Error(respErr.Error())
		return nil, false
	}

	bodyStr, _ := ioutil.ReadAll(resp.Body)
	if _vari.Verbose {
		_logUtils.PrintUnicode(bodyStr)
	}

	var result _domain.RpcResult
	json.Unmarshal(bodyStr, &result)

	defer resp.Body.Close()

	code := result.Code

	return result, code == _const.ResultSuccess
}

func PostRasa(url string, params interface{}) (ret _domain.RpcResult, success bool) {
	if _vari.Verbose {
		_logUtils.Info(url)
	}
	client := &http.Client{}

	paramStr, err := json.Marshal(params)
	if err != nil {
		_logUtils.Error(err.Error())
		return
	}

	req, reqErr := http.NewRequest("POST", url, strings.NewReader(string(paramStr)))
	if reqErr != nil {
		_logUtils.Error(reqErr.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, respErr := client.Do(req)
	if respErr != nil {
		_logUtils.Error(respErr.Error())
		return
	}

	bodyStr, _ := ioutil.ReadAll(resp.Body)
	if _vari.Verbose {
		_logUtils.PrintUnicode(bodyStr)
	}
	defer resp.Body.Close()

	var rasaResp domain.RasaResp
	json.Unmarshal(bodyStr, &rasaResp)

	success = rasaResp.Intent.ID != 0
	ret.Payload = rasaResp

	return
}

func GenUrl(server string, path string) string {
	server = UpdateUrl(server)
	url := fmt.Sprintf("%sapi/v1/%s", server, path)
	return url
}

func UpdateUrl(url string) string {
	if strings.LastIndex(url, "/") < len(url)-1 {
		url += "/"
	}
	return url
}
