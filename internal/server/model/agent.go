package model

import (
	consts "github.com/utlai/utl/internal/comm/const"
	"time"
)

type Agent struct {
	BaseModel

	Name             string             `json:"name"`
	Desc             string             `json:"desc"`
	Status           consts.AgentStatus `json:"status"`
	LastRegisterTime time.Time          `json:"lastRegisterTime"`

	Ip         string `json:"ip"`
	Port       int    `json:"port"`
	MacAddress string `json:"macAddress"`
	RpcPort    int    `json:"rpcPort"`
	SshPort    int    `json:"sshPort"`
	VncAddress string `json:"vncAddress"`
	WorkDir    string `json:"workDir"`
}

func (Agent) TableName() string {
	return "biz_agent"
}
