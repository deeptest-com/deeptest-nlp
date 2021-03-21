package agentModel

import _const "github.com/utlai/utl/internal/pkg/const"

type Config struct {
	Server   string              `yaml:"Server"`
	Ip       string              `yaml:"ip"`
	Port     int                 `yaml:"port"`
	Platform _const.WorkPlatform `yaml:"platform"`

	Language string
	HostName string
	WorkDir  string
	LogDir   string
}
