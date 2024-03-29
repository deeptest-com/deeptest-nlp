// +build test perm api

package tests

import (
	"fmt"
	"github.com/utlai/utl/cmd/server/tests/mock"
	"github.com/bxcodec/faker/v3"
	"github.com/fatih/color"
	"testing"

	"github.com/kataras/iris/v12"
)

func TestPermissions(t *testing.T) {
	getMore(t, "permissions", iris.StatusOK, 200, "操作成功")
}

func TestPermissionCreate(t *testing.T) {
	m := mock.Permission{}
	err := faker.FakeData(&m)
	if err != nil {
		color.Red("TestPermissionCreate %+v", err)
		return
	}

	create(t, "permissions", m, iris.StatusOK, 200, "操作成功")
}

func TestPermissionUpdate(t *testing.T) {
	m := mock.Permission{}
	err := faker.FakeData(&m)
	if err != nil {
		color.Red("TestPermissionUpdate %+v", err)
		return
	}

	tr, err := CreatePermission()
	if err != nil {
		fmt.Print(err)
	}

	url := "permissions/%d"
	update(t, fmt.Sprintf(url, tr.ID), m, iris.StatusOK, 200, "操作成功")
}

func TestPermissionDelete(t *testing.T) {
	tr, err := CreatePermission()
	if err != nil {
		fmt.Print(err)
	}
	delete(t, fmt.Sprintf("permissions/%d", tr.ID), iris.StatusOK, 200, "删除成功")
}
