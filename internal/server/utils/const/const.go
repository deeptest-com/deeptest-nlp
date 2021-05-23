package serverConst

type VmPlatform string

const (
	AppName    = "server"
	NluVersion = "2.0"

	PageSize            = 15
	Kvm      VmPlatform = "kvm"
	Pve      VmPlatform = "pve"

	Docker    ContainerPlatform = "docker"
	Portainer ContainerPlatform = "portainer"
)

var (
	SlotTypeAbbrMap = map[string]string{"synonym": "syn", "lookup": "lkp", "regex": "rgx"}
)

type ContainerPlatform string
