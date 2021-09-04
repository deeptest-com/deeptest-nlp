package agentModel

type Config struct {
	Server     string `json:"server" yaml:"Server"`
	Ip         string `json:"ip" yaml:"ip"`
	Port       int    `json:"port" yaml:"port"`
	MacAddress string `json:"macAddress" yaml:"macAddress"`

	Language string
	HostName string
	WorkDir  string
	LogDir   string
}
