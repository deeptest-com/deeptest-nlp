package domain

import consts "github.com/utlai/utl/internal/comm/const"

type Agent struct {
	MacAddress string `json:"macAddress"`
	Ip         string `json:"ip"`
	Port       int    `json:"port"`
	WorkDir    string `json:"workDir"`

	Status consts.AgentStatus `json:"status"`
}
