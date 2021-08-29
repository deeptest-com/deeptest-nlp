package _const

import "os"

const (
	PthSep = string(os.PathSeparator)

	UserTokenExpireTime = 365 * 24 * 60 * 60 * 1000

	RetryTime    = 3
	AgentRunTime = 20 // sec

	WebCheckQueueInterval    = 10 // sec
	AgentCheckDeviceInterval = 10 // sec

	MaxVmOnHost = 3
	RpcPort     = 8972

	UploadDir = "uploads"

	PageSize = 10
)
