package net

import (
	"encoding/json"
	consts "github.com/utlai/utl/internal/comm/const"
	"github.com/utlai/utl/internal/comm/domain"
	_fileUtils "github.com/utlai/utl/internal/pkg/libs/file"
	_httpUtils "github.com/utlai/utl/internal/pkg/libs/http"
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"testing"
)

func TestUpload(t *testing.T) {
	_logUtils.Init(consts.AppNameAgent)
	result := domain.TestResult{Name: "RasaResult Name"}

	zipFile := "/Users/aaron/testResult.zip"

	result.Payload = nil
	uploadResultUrl := _httpUtils.GenUrl("http://localhost:8085/", "client/build/uploadResult")

	files := []string{zipFile}
	extraParams := map[string]string{}
	json, _ := json.Marshal(result)
	extraParams["result"] = string(json)

	_fileUtils.Upload(uploadResultUrl, files, extraParams)
}
